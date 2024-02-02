package db

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/internal/logging"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

//go:generate mockgen -destination=./individual_mock.go -package=db . IndividualRepo

type IndividualRepo interface {
	GetAll(ctx context.Context, options api.ListIndividualsOptions) ([]*api.Individual, error)
	GetByID(ctx context.Context, id string) (*api.Individual, error)
	Put(ctx context.Context, individual *api.Individual, fields containers.StringSet) (*api.Individual, error)
	PutMany(ctx context.Context, individuals []*api.Individual, fields containers.StringSet) ([]*api.Individual, error)
	PerformAction(ctx context.Context, id string, action IndividualAction) error
	PerformActionMany(ctx context.Context, ids containers.StringSet, action IndividualAction) error
	FindDuplicates(ctx context.Context, df dataframe.DataFrame, deduplicationConfig deduplication.DeduplicationConfig) ([]*api.Individual, error)
}

type individualRepo struct {
	db *sqlx.DB
}

func (i individualRepo) FindDuplicates(ctx context.Context, df dataframe.DataFrame, deduplicationConfig deduplication.DeduplicationConfig) ([]*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		return i.findDuplicatesInternal(ctx, tx, df, deduplicationConfig)
	})
	if err != nil {
		return nil, err
	}
	return ret.([]*api.Individual), nil
}

func (i individualRepo) findDuplicatesInternal(ctx context.Context, tx *sqlx.Tx, df dataframe.DataFrame, config deduplication.DeduplicationConfig) ([]*api.Individual, error) {
	if i.driverName() != "postgres" {
		return nil, fmt.Errorf("deduplication is only implemented for postgres")
	}

	ret := make([]*api.Individual, 0)

	selectedCountryID, err := utils.GetSelectedCountryID(ctx)

	// first we get the schema of the individual_registrations table
	var schema []DBColumn
	schemaQuery := buildTableSchemaQuery()
	err = tx.SelectContext(ctx, &schema, schemaQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get table schema: %w", err)
	}

	// then we collect the columns that are relevant for deduplication
	columnsOfInterest := []string{}
	uploadDfHasIdColumn := slices.Contains(df.Names(), constants.DBColumnIndividualID)
	if uploadDfHasIdColumn {
		columnsOfInterest = append(columnsOfInterest, constants.DBColumnIndividualID)
	}
	for _, d := range config.Types {
		cols := d.Config.Columns
		for c := range cols {
			columnsOfInterest = append(columnsOfInterest, cols[c])
		}
	}

	// we create a temp table with the relevant columns
	// we use the request id to make sure the temp table name is unique
	tempTableName := strings.Replace(fmt.Sprintf("temp_individuals_%s", utils.GetRequestID(ctx)), "-", "_", -1)
	createTempTableQuery := buildCreateTempTableQuery(tempTableName, schema, columnsOfInterest)
	result := tx.MustExec(createTempTableQuery)
	if result == nil {
		return nil, fmt.Errorf("failed to create temp table")
	}

	// we insert the data from the upload into the temp table
	df = df.Select(columnsOfInterest)
	insertQuery, args := buildInsertIndividualsQuery(tempTableName, schema, df, columnsOfInterest, uploadDfHasIdColumn)
	result = tx.MustExec(insertQuery, args...)
	if result == nil {
		return nil, fmt.Errorf("failed to insert into temp table")
	}

	// now we look for duplicates in a cross join between the temp table and the individual_registrations table
	deduplicationQuery := buildDeduplicationQuery(tempTableName, columnsOfInterest, config, uploadDfHasIdColumn, schema)
	err = tx.SelectContext(ctx, &ret, deduplicationQuery, selectedCountryID)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type DBColumn struct {
	Name    string
	Default *string
	SQLType string
}

func buildTableSchemaQuery() string {
	b := &strings.Builder{}
	b.WriteString("SELECT column_name as Name, udt_name as SQLType, column_default as Default FROM information_schema.columns WHERE table_name = 'individual_registrations';")
	return b.String()
}

/*
EXAMPLE:

CREATE TEMPORARY TABLE temp_individuals_<requestId>

	(birth_date date,first_name varchar,middle_name varchar,last_name varchar,native_name varchar,email_1 varchar,email_2 varchar,email_3 varchar)
	ON COMMIT DROP;
*/
func buildCreateTempTableQuery(tempTableName string, schema []DBColumn, columnsOfInterest []string) string {
	b := &strings.Builder{}

	b.WriteString(fmt.Sprintf("CREATE TEMPORARY TABLE %s", tempTableName))
	columns := []string{}
	for si, _ := range schema {
		if slices.Contains(columnsOfInterest, schema[si].Name) {
			columns = append(columns, fmt.Sprintf("%s %s", schema[si].Name, schema[si].SQLType))
		}
	}
	b.WriteString(fmt.Sprintf(" (%s)", strings.Join(columns, ",")))
	b.WriteString(" ON COMMIT DROP;")
	return b.String()
}

/*
EXAMPLE:

INSERT INTO temp_individuals_<requestId>

	SELECT * FROM UNNEST($1::date[],$2::varchar[],$3::varchar[],$4::varchar[],$5::varchar[],$6::varchar[],$7::varchar[],$8::varchar[]);

where the parameters are the values of the relevant columns in the dataframe df
*/
func buildInsertIndividualsQuery(tempTableName string, schema []DBColumn, df dataframe.DataFrame, columnsOfInterest []string, uploadDfHasIdColumn bool) (string, []interface{}) {
	b := &strings.Builder{}

	b.WriteString(fmt.Sprintf("INSERT INTO %s SELECT * FROM UNNEST(", tempTableName))

	var args []interface{}
	var types []string
	for _, col := range schema {
		// we only insert the columns that are relevant for deduplication and the id column if it exists in the upload
		if slices.Contains(columnsOfInterest, col.Name) || (uploadDfHasIdColumn && col.Name == constants.DBColumnIndividualID) {
			g := df.Col(col.Name)
			if g.Err != nil {
				args = append(args, pq.Array(nil))
			} else {
				if col.SQLType == "date" {
					values := make([]*time.Time, df.Nrow())
					for i := 0; i < df.Nrow(); i++ {
						t, err := time.Parse("2006-01-02", g.Elem(i).String())
						if err != nil {
							values[i] = nil
							continue
						}
						values[i] = &t
					}
					args = append(args, pq.Array(values))
				} else if col.SQLType == "uuid" {
					values := make([]uuid.UUID, df.Nrow())
					for i := 0; i < df.Nrow(); i++ {
						v, err := uuid.Parse(g.Elem(i).String())
						if err != nil {
							values[i] = uuid.Nil
							continue
						}
						values[i] = v
					}
					args = append(args, pq.Array(values))
				} else {
					args = append(args, pq.Array(g.Records()))
				}
			}
			types = append(types, fmt.Sprintf("$%d::%s[]", len(args), col.SQLType))
		}
	}

	b.WriteString(strings.Join(types, ","))
	b.WriteString(");")
	return b.String(), args
}

/*
EXAMPLE:

SELECT DISTINCT  ir.birth_date,ir.email_1,ir.email_2,ir.email_3,ir.first_name,ir.middle_name,ir.last_name,ir.native_name,ir.id

	FROM individual_registrations ir
	CROSS JOIN temp_individuals_<requestId> ti
	WHERE ir.country_id = $1
		AND ir.deleted_at IS NULL
		AND (ti.birth_date = ir.birth_date)
		AND (ti.email_1 = ir.email_1 OR ti.email_2 = ir.email_2 OR ti.email_3 = ir.email_3)
		AND (ti.first_name = ir.first_name AND ti.middle_name = ir.middle_name AND ti.last_name = ir.last_name AND ti.native_name = ir.native_name);
*/
func buildDeduplicationQuery(tempTableName string, columnsOfInterest []string, config deduplication.DeduplicationConfig, uploadDfHasIdColumn bool, schema []DBColumn) string {
	b := &strings.Builder{}

	b.WriteString("SELECT DISTINCT ")
	for c := range columnsOfInterest {
		b.WriteString(fmt.Sprintf("ir.%s,", columnsOfInterest[c]))
	}
	b.WriteString("ir.id FROM individual_registrations ir")
	b.WriteString(fmt.Sprintf(" CROSS JOIN %s ti", tempTableName))
	b.WriteString(" WHERE ir.country_id = $1 AND ir.deleted_at IS NULL")

	subQueries := []string{}
	for _, dt := range config.Types {
		subSubQueries := []string{}
		cols := dt.Config.Columns
		for _, c := range cols {
			subSubQueries = append(subSubQueries, fmt.Sprintf("ti.%s = ir.%s", c, c))
		}
		subQueries = append(subQueries, strings.Join(subSubQueries, fmt.Sprintf(" %s ", dt.Config.Condition)))
	}
	if len(subQueries) > 0 {
		b.WriteString(fmt.Sprintf(" AND (%s)", strings.Join(subQueries, fmt.Sprintf(") %s (", config.Operator))))
	}
	if uploadDfHasIdColumn {
		b.WriteString(" AND ti.id::uuid NOT IN (SELECT id FROM individual_registrations)")
	}
	b.WriteString(";")

	return b.String()
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
		return nil, errors.New(locales.GetTranslator()("error_unexpected_number", len(ret)))
	}
	return ret[0], nil
}

func (i individualRepo) PerformAction(ctx context.Context, id string, action IndividualAction) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.performActionManyInternal(ctx, tx, containers.NewStringSet(id), action)
		return nil, err
	})
	return err
}

func (i individualRepo) PerformActionMany(ctx context.Context, ids containers.StringSet, action IndividualAction) error {
	_, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		err := i.performActionManyInternal(ctx, tx, ids, action)
		return nil, err
	})
	return err
}

func (i individualRepo) performActionManyInternal(ctx context.Context, tx *sqlx.Tx, ids containers.StringSet, action IndividualAction) error {
	l := logging.NewLogger(ctx).With(zap.Strings("individual_ids", ids.Items()))
	l.Debug("performing action: " + string(action) + " individuals")

	if err := batch(maxParams/ids.Len(), ids.Items(), func(idsInBatch []string) (bool, error) {
		var query string

		switch action {
		case DeleteAction:
			query = "DELETE FROM individual_registrations WHERE id IN ("
		case ActivateAction:
			query = "UPDATE individual_registrations SET inactive = false WHERE id IN ("
		case DeactivateAction:
			query = "UPDATE individual_registrations SET inactive = true WHERE id IN ("
		}

		var args []interface{}
		for i, id := range idsInBatch {
			if i != 0 {
				query += ","
			}
			query += fmt.Sprintf("$%d", i+1)
			args = append(args, id)
		}
		query += ") "

		auditDuration := logDuration(ctx, "performing action: "+string(action)+" individuals", zap.Int("count", len(idsInBatch)))
		defer auditDuration()

		result, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			l.Error("failed to "+string(action)+" individuals", zap.Error(err))
			return false, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			l.Error("failed to get rows affected", zap.Error(err))
			return false, err
		} else if rowsAffected != int64(len(idsInBatch)) {
			l.Error("failed to "+string(action)+" all individuals", zap.Int64("rows_affected", rowsAffected))
			return false, fmt.Errorf("failed to " + string(action) + " all individuals")
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
