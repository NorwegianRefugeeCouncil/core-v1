package db

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"go.uber.org/zap"
)

//go:generate mockgen -destination=./individual_mock.go -package=db . IndividualRepo

type IndividualRepo interface {
	GetAll(ctx context.Context, options api.ListIndividualsOptions) ([]*api.Individual, error)
	GetByID(ctx context.Context, id string) (*api.Individual, error)
	Put(ctx context.Context, individual *api.Individual, fields containers.StringSet) (*api.Individual, error)
	PutMany(ctx context.Context, individuals []*api.Individual, fields containers.StringSet) ([]*api.Individual, error)
	PerformAction(ctx context.Context, id string, action string) error
	PerformActionMany(ctx context.Context, ids containers.StringSet, action string) error
	FindDuplicates(ctx context.Context, individuals []*api.Individual, deduplicationTypes []deduplication.DeduplicationTypeName) ([]*api.Individual, error)
}

type individualRepo struct {
	db *sqlx.DB
}

func (i individualRepo) FindDuplicates(ctx context.Context, individuals []*api.Individual, deduplicationTypes []deduplication.DeduplicationTypeName) ([]*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.findDuplicatesInternal(ctx, tx, individuals, deduplicationTypes)
	})
	if err != nil {
		return nil, err
	}
	return ret.([]*api.Individual), nil
}

func (i individualRepo) findDuplicatesInternal(ctx context.Context, tx *sqlx.Tx, newIndividuals []*api.Individual, deduplicationTypes []deduplication.DeduplicationTypeName) ([]*api.Individual, error) {
	if i.driverName() != "postgres" {
		return nil, fmt.Errorf("deduplication is only implemented for postgres")
	}

	args := make([]interface{}, 0)
	ret := make([]*api.Individual, 0)

	selectedCountryID, err := utils.GetSelectedCountryID(ctx)
	if err != nil {
		return nil, err
	}
	query, args := buildDeduplicationQuery(selectedCountryID, newIndividuals, deduplicationTypes)
	out := make([]*api.Individual, 0)

	err = tx.SelectContext(ctx, &out, query, args...)
	if err != nil {
		return nil, err
	}
	ret = append(ret, out...)
	return ret, nil
}

type columnValues = map[string][]string
type queryValues = map[deduplication.DeduplicationTypeName]columnValues

/*
example of queryValues:
{
  "IDs": {
    "id_1": ["ID1","ID2","ID3","ID4","ID5","ID6"],
    "id_2": ["ID1","ID2","ID3","ID4","ID5","ID6"],
    "id_3": ["ID1","ID2","ID3","ID4","ID5","ID6"],
  },
  "Names": {
    "firstName": ["John","Jane"],
    "lastName": ["Doe","Doe"]
  }
}
*/

func buildDeduplicationQuery(selectedCountryID string, newIndividuals []*api.Individual, deduplicationTypes []deduplication.DeduplicationTypeName) (string, []interface{}) {
	b := &strings.Builder{}
	b.WriteString("SELECT * FROM individual_registrations WHERE ")
	b.WriteString("country_id = '")
	b.WriteString(selectedCountryID)
	b.WriteString("' AND deleted_at IS NULL")

	params := collectParams(newIndividuals, deduplicationTypes)
	args := fillQueryWithParameters(params, b)
	return b.String(), args
}

// building a map of values and parameters first to avoid empty values
func collectParams(individuals []*api.Individual, deduplicationTypes []deduplication.DeduplicationTypeName) queryValues {
	params := queryValues{}
	for d := 0; d < len(deduplicationTypes); d++ {
		valuesPerColumn := columnValues{}
		deduplicationType := deduplicationTypes[d]
		deduplicationConfig := deduplication.DeduplicationTypes[deduplicationType].Config
		if deduplicationConfig.Condition == deduplication.LOGICAL_OPERATOR_OR {
			collectParamsForOrQuery(individuals, deduplicationConfig, valuesPerColumn)
		} else if deduplicationConfig.Condition == deduplication.LOGICAL_OPERATOR_AND {
			collectParamsForAndQuery(individuals, deduplicationConfig, valuesPerColumn)
		}

		if len(valuesPerColumn) > 0 {
			params[deduplicationType] = valuesPerColumn
		}
	}
	return params
}

func collectParamsForOrQuery(individuals []*api.Individual, deduplicationConfig deduplication.DeduplicationTypeValue, allValuesPerColumns columnValues) columnValues {
	values := make([]string, 0, len(individuals))
	for c := 0; c < len(deduplicationConfig.Columns); c++ {
		column := deduplicationConfig.Columns[c]
		for _, individual := range individuals {
			v, err := individual.GetFieldValue(column)
			if err == nil && v != "" && v != nil {
				values = append(values, v.(string))
			}
		}
	}
	if len(values) > 0 {
		for c := 0; c < len(deduplicationConfig.Columns); c++ {
			column := deduplicationConfig.Columns[c]
			allValuesPerColumns[column] = values
		}
	}
	return allValuesPerColumns
}

func collectParamsForAndQuery(individuals []*api.Individual, deduplicationConfig deduplication.DeduplicationTypeValue, allValuesPerColumns columnValues) columnValues {

	for c := 0; c < len(deduplicationConfig.Columns); c++ {
		values := make([]string, 0, len(individuals))
		column := deduplicationConfig.Columns[c]
		for _, individual := range individuals {
			v, err := individual.GetFieldValue(column)
			if err == nil && v != "" && v != nil {
				values = append(values, v.(string))
			}
		}
		if len(values) > 0 {
			allValuesPerColumns[column] = values
		}
	}
	return allValuesPerColumns
}

func fillQueryWithParameters(paramMap queryValues, b *strings.Builder) []interface{} {
	args := make([]interface{}, 0)

	for typeKey, typeValues := range paramMap {
		b.WriteString(" AND (")
		queryParts := make([]string, 0)
		for colKey, colValues := range typeValues {
			args = append(args, pq.Array(colValues))
			queryParts = append(queryParts, fmt.Sprintf("%s IN (SELECT * FROM UNNEST ($%d::text[]))", colKey, len(args)))
		}
		b.WriteString(strings.Join(queryParts, " "+deduplication.DeduplicationTypes[typeKey].Config.Condition+" "))
		b.WriteString(")")
	}

	return args
}

func (i individualRepo) driverName() string {
	d := i.db.DriverName()
	if d == "sqlite3" {
		return "sqlite"
	} else if d == "postgres" {
		return "postgres"
	} else {
		panic("unsupported driver")
	}
}

func (i individualRepo) GetAll(ctx context.Context, options api.ListIndividualsOptions) ([]*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.batchedGetAllInternal(ctx, tx, options)
	})
	if err != nil {
		return nil, err
	}
	return ret.([]*api.Individual), nil
}

func (i individualRepo) batchedGetAllInternal(ctx context.Context, tx *sqlx.Tx, options api.ListIndividualsOptions) ([]*api.Individual, error) {
	if options.IDs.Len() > 0 {
		var ret []*api.Individual
		err := batch(maxParams/options.IDs.Len(), options.IDs.Items(), func(idsInBatch []string) (stop bool, err error) {
			// check if we already reached the limit. If so, exit early
			if options.Take != 0 && len(ret) >= options.Take {
				return true, nil
			}
			optionsForBatch := options
			optionsForBatch.IDs = containers.NewStringSet(idsInBatch...)
			optionsForBatch.Take = utils.Max(0, options.Take-len(ret))
			individualsInBatch, err := i.unbatchedGetAllInternal(ctx, tx, optionsForBatch)
			if err != nil {
				return false, err
			}
			ret = append(ret, individualsInBatch...)
			return false, nil
		})
		if err != nil {
			return nil, err
		}
		// todo: batch sorting. Because we're batching, we can't rely on the database to do the sorting between batches.
		// The records within a batch are sorted, but they need to be sorted across batches.
		//
		// When performing a query, the max number of parameters for postgres is 65535. If we only filter by
		// ids, we can have up to ~65535 ids per batch. This is unlikely to
		// happen for the webserver usecase. Though, when integrations or internal services will use this API,
		// this could be a problem.
		//
		// Option 1: sort in-memory:
		//   we need to assemble the batches in-memory first, then sort them with the same sorting
		//   algorithms as the database.
		//
		// Option 2: temporary table:
		//   copy the batch results into a temporary table, then query this temporary table. This is probably
		//   the best and simplest option, as we don't need to implement the sorting algorithm in go.
		return ret, nil
	}
	return i.unbatchedGetAllInternal(ctx, tx, options)
}

func (i individualRepo) unbatchedGetAllInternal(ctx context.Context, tx *sqlx.Tx, options api.ListIndividualsOptions) ([]*api.Individual, error) {
	l := logging.NewLogger(ctx)
	l.Debug("getting list individuals", zap.Any("options", options))
	var ret []*api.Individual

	auditDuration := logDuration(ctx, "get individuals unbatched")
	defer auditDuration()

	sql, args := newGetAllIndividualsSQLQuery(i.driverName(), options).build()
	err := tx.SelectContext(ctx, &ret, sql, args...)
	if err != nil {
		l.Error("failed to list individuals", zap.Error(err))
	}
	return ret, nil
}

func (i individualRepo) GetByID(ctx context.Context, id string) (*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.getByIdInternal(ctx, tx, id)
	})
	if err != nil {
		return nil, err
	}
	return ret.(*api.Individual), nil
}

func (i individualRepo) getByIdInternal(ctx context.Context, tx *sqlx.Tx, id string) (*api.Individual, error) {
	l := logging.NewLogger(ctx).With(zap.String("individual_id", id))
	l.Debug("getting individual by id")

	auditDuration := logDuration(ctx, "get individual by id")
	defer auditDuration()

	var ret = api.Individual{}
	err := tx.GetContext(ctx, &ret, "SELECT * FROM individual_registrations WHERE id = $1 and deleted_at IS NULL", id)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (i individualRepo) PutMany(ctx context.Context, individuals []*api.Individual, fields containers.StringSet) ([]*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.putManyInternal(ctx, tx, individuals, fields)
	})
	if err != nil {
		return nil, err
	}
	return ret.([]*api.Individual), nil
}

func (i individualRepo) putManyInternal(ctx context.Context, tx *sqlx.Tx, individuals []*api.Individual, fields containers.StringSet) ([]*api.Individual, error) {

	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)

	fieldsSet := fields.Clone()
	if fieldsSet.Contains("phone_number_1") {
		fieldsSet.Add("normalized_phone_number_1")
	}
	if fieldsSet.Contains("phone_number_2") {
		fieldsSet.Add("normalized_phone_number_2")
	}
	if fieldsSet.Contains("phone_number_3") {
		fieldsSet.Add("normalized_phone_number_3")
	}
	fieldsSet.Remove("deleted_at")
	fieldsSet.Remove("created_at")
	fieldsSet.Remove("updated_at")
	fieldsSet.Add("id")

	fieldSlice := fieldsSet.Items()

	ret := make([]*api.Individual, 0, len(individuals))
	if err := batch(maxParams/len(fieldSlice), individuals, func(individualsInBatch []*api.Individual) (bool, error) {
		args := make([]interface{}, 0)
		b := &strings.Builder{}
		b.WriteString("INSERT INTO individual_registrations (" + strings.Join(fieldSlice, ",") + ",created_at,updated_at) VALUES ")

		for i, individual := range individualsInBatch {
			if individual.ID == "" {
				individual.ID = uuid.New().String()
			}
			if i != 0 {
				b.WriteString(",")
			}
			b.WriteString("(")
			for j, field := range fieldSlice {
				if j != 0 {
					b.WriteString(",")
				}
				fieldValue, err := individual.GetFieldValue(field)
				if err != nil {
					return false, err
				}
				args = append(args, fieldValue)
				b.WriteString(fmt.Sprintf("$%d", len(args)))
			}
			b.WriteString(",'" + nowStr + "','" + nowStr + "')")
		}
		b.WriteString(" ON CONFLICT (id) DO UPDATE SET ")
		isFirst := true
		for _, field := range fieldSlice {
			if field == "id" || field == "country_id" {
				continue
			}
			if !isFirst {
				b.WriteString(",")
			}
			isFirst = false
			b.WriteString(fmt.Sprintf("%s = EXCLUDED.%s", field, field))
		}
		b.WriteString(", updated_at='" + nowStr + "'")
		b.WriteString(" RETURNING *")

		var out []*api.Individual
		qry := b.String()

		auditDuration := logDuration(ctx, "putting individuals", zap.Int("count", len(individualsInBatch)))
		defer auditDuration()

		err := tx.SelectContext(ctx, &out, qry, args...)
		if err != nil {
			return false, err
		}
		ret = append(ret, out...)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (i individualRepo) Put(ctx context.Context, individual *api.Individual, fields containers.StringSet) (*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.putInternal(ctx, tx, individual, fields)
	})
	if err != nil {
		return nil, err
	}
	return ret.(*api.Individual), nil
}

func (i individualRepo) putInternal(ctx context.Context, tx *sqlx.Tx, individual *api.Individual, fields containers.StringSet) (*api.Individual, error) {
	ret, err := i.putManyInternal(ctx, tx, []*api.Individual{individual}, fields)
	if err != nil {
		return nil, err
	}
	if len(ret) != 1 {
		return nil, fmt.Errorf("unexpected number of individuals returned: %d", len(ret))
	}
	return ret[0], nil
}

func (i individualRepo) PerformAction(ctx context.Context, id string, action string) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.performActionManyInternal(ctx, tx, containers.NewStringSet(id), action)
		return nil, err
	})
	return err
}

func (i individualRepo) PerformActionMany(ctx context.Context, ids containers.StringSet, action string) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.performActionManyInternal(ctx, tx, ids, action)
		return nil, err
	})
	return err
}

func (i individualRepo) performActionManyInternal(ctx context.Context, tx *sqlx.Tx, ids containers.StringSet, action string) error {
	l := logging.NewLogger(ctx).With(zap.Strings("individual_ids", ids.Items()))
	l.Debug("performing action: " + action + " individuals")

	if err := batch(maxParams/ids.Len(), ids.Items(), func(idsInBatch []string) (bool, error) {
		var query = "UPDATE individual_registrations SET " + individualActions[action].targetField + " = $1 WHERE id IN ("
		var args = []interface{}{individualActions[action].newValue}
		for i, id := range idsInBatch {
			if i != 0 {
				query += ","
			}
			query += fmt.Sprintf("$%d", i+2)
			args = append(args, id)
		}
		query += ") "
		for _, c := range individualActions[action].conditions {
			query += c + " "
		}

		auditDuration := logDuration(ctx, "performing action: "+action+" individuals", zap.Int("count", len(idsInBatch)))
		defer auditDuration()

		result, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			l.Error("failed to "+action+" individuals", zap.Error(err))
			return false, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			l.Error("failed to get rows affected", zap.Error(err))
			return false, err
		} else if rowsAffected != int64(len(idsInBatch)) {
			l.Error("failed to "+action+" all individuals", zap.Int64("rows_affected", rowsAffected))
			return false, fmt.Errorf("failed to " + action + " all individuals")
		}

		return false, nil
	}); err != nil {
		return err
	}

	return nil
}

func NewIndividualRepo(db *sqlx.DB) IndividualRepo {
	return &individualRepo{db: db}
}
