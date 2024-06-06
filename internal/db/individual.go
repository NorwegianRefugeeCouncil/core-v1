package db

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/locales"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"golang.org/x/exp/slices"

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
	FindDuplicates(ctx context.Context, individuals []*api.Individual, deduplicationConfig deduplication.DeduplicationConfig) ([]containers.Set[int], map[int][]*api.Individual, error) 
}

type individualRepo struct {
	db *sqlx.DB
}

type fileDuplicateRet struct {
	IdxA int `db:"idxa"`
	IdxB int `db:"idxb"`
}

type dbDuplicateRet struct {
	api.Individual
	Idx int `db:"idx"`
}

func (i individualRepo) FindDuplicates(ctx context.Context, individuals []*api.Individual, deduplicationConfig deduplication.DeduplicationConfig) ([]containers.Set[int], map[int][]*api.Individual, error) {
	ret, err := doInTransaction(ctx, i.db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		fileDuplicates, dbDuplicates, err := i.findDuplicatesInternal(ctx, tx, individuals, deduplicationConfig)
		return []interface{}{fileDuplicates, dbDuplicates}, err
	})
	if err != nil {
		return nil, nil, err
	}
	return ret.([]interface{})[0].([]containers.Set[int]), ret.([]interface{})[1].(map[int][]*api.Individual), nil 
}

func (i individualRepo) findDuplicatesInternal(ctx context.Context, tx *sqlx.Tx, individuals []*api.Individual, config deduplication.DeduplicationConfig) ([]containers.Set[int], map[int][]*api.Individual, error) {
	if i.driverName() != "postgres" {
		return nil, nil, fmt.Errorf("deduplication is only implemented for postgres")
	}

	selectedCountryID, err := utils.GetSelectedCountryID(ctx)
	if err != nil {
		return nil, nil, err
	}

	deduplicationTempTableConfig, err := prepareDeduplicationTempTable(ctx, tx, individuals, config)
	if err != nil {
		return nil, nil, err
	}

	fileDuplicates, err := findFileDuplicates(ctx, tx, deduplicationTempTableConfig, config, len(individuals))
	if err != nil {
		return nil, nil, err
	}
	for _, d := range fileDuplicates {
		if d.Len() > 0 {
			return fileDuplicates, nil, nil
		}
	}

	dbDuplicates, err := findDbDuplicates(ctx, tx, deduplicationTempTableConfig, config, selectedCountryID)
	if err != nil {
		return nil, nil, err
	}
	if len(dbDuplicates) > 0 {
		return nil, dbDuplicates, nil
	}

	return nil, nil, nil	
}

func findDbDuplicates(ctx context.Context, tx *sqlx.Tx, deduplicationTempTableConfig *DeduplicationTempTableConfig, config deduplication.DeduplicationConfig, selectedCountryID string) (map[int][]*api.Individual, error) {
	ret := make([]*dbDuplicateRet, 0)

	// now we look for duplicates in a cross join between the temp table and the individual_registrations table
	deduplicationQuery := buildDbDeduplicationQuery(
		deduplicationTempTableConfig.tempTableName,
		deduplicationTempTableConfig.columnsOfInterest,
		config,
		deduplicationTempTableConfig.schema,
	)
	err := tx.SelectContext(ctx, &ret, deduplicationQuery, selectedCountryID)
	if err != nil {
		return nil, err
	}

	duplicates := make(map[int][]*api.Individual)
	for _, d := range ret {
		idx := d.Idx -1
		if (duplicates[idx] == nil) {
			duplicates[idx] = make([]*api.Individual, 0)
		}
		duplicates[idx] = append(duplicates[idx], &d.Individual)
	}

	return duplicates, nil
}

func findFileDuplicates(ctx context.Context, tx *sqlx.Tx, deduplicationTempTableConfig *DeduplicationTempTableConfig, config deduplication.DeduplicationConfig, individualCount int) ([]containers.Set[int], error) {
	ret := make([]*fileDuplicateRet, 0)

	// now we look for duplicates in a cross join between the temp table and the individual_registrations table
	deduplicationQuery := buildFileDeduplicationQuery(
		deduplicationTempTableConfig.tempTableName,
		deduplicationTempTableConfig.columnsOfInterest,
		config,
		deduplicationTempTableConfig.schema,
	)
	err := tx.SelectContext(ctx, &ret, deduplicationQuery)
	if err != nil {
		return nil, err
	}

	duplicateScores := []containers.Set[int]{}
	for i := 0; i < individualCount; i++ {
		duplicateScores = append(duplicateScores, containers.NewSet[int]())
	}

	for _, d := range ret {
		duplicateScores[d.IdxA - 1].Add(d.IdxB - 1)
	}

	// The query returns duplicates in both directions, so we need to remove duplicates that are not mutual
	for i := range(duplicateScores) {
		for j := range(duplicateScores[i]) {
			if !duplicateScores[j].Contains(i) {
				duplicateScores[i].Remove(j)
			}
		}
	}

	return duplicateScores, nil
}

type DBColumn struct {
	Name    string
	Default *string
	SQLType string
}

type DeduplicationTempTableConfig struct {
	tempTableName string
	columnsOfInterest []string
	schema []DBColumn
}

func prepareDeduplicationTempTable(ctx context.Context, tx *sqlx.Tx, individuals []*api.Individual, config deduplication.DeduplicationConfig) (deduplicationTempTableConfig *DeduplicationTempTableConfig, error error) {
	// first we get the schema of the individual_registrations table
	var schema []DBColumn
	schemaQuery := buildTableSchemaQuery()
	err := tx.SelectContext(ctx, &schema, schemaQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get table schema: %w", err)
	}

	// then we collect the columns that are relevant for deduplication
	columnsOfInterest := []string{}
	columnsOfInterest = append(columnsOfInterest, constants.DBColumnIndividualID)
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
	err = insertTempIndividuals(ctx, tx, tempTableName, schema, individuals, columnsOfInterest)
	if err != nil { 
		return nil, fmt.Errorf("failed to insert into temp table")
	}

	return &DeduplicationTempTableConfig{
		tempTableName: tempTableName,
		columnsOfInterest: columnsOfInterest,
		schema: schema,
	}, nil
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
	columns = append(columns, "idx serial")
	b.WriteString(fmt.Sprintf(" (%s)", strings.Join(columns, ",")))
	b.WriteString(" ON COMMIT DROP;")

	// add index for each column
	for si, _ := range schema {
		if slices.Contains(columnsOfInterest, schema[si].Name) {
			b.WriteString(fmt.Sprintf("CREATE INDEX ON %s (%s);", tempTableName, schema[si].Name))
		}
	}
	b.WriteString("CREATE INDEX ON " + tempTableName + " (idx);")

	return b.String()
}

/*
EXAMPLE:

INSERT INTO temp_individuals_<requestId>

	SELECT * FROM UNNEST($1::date[],$2::varchar[],$3::varchar[],$4::varchar[],$5::varchar[],$6::varchar[],$7::varchar[],$8::varchar[]);

where the parameters are the values of the relevant columns in the dataframe df
*/
func insertTempIndividuals(ctx context.Context, tx *sqlx.Tx, tempTableName string, schema []DBColumn, individuals []*api.Individual, columnsOfInterest []string) error {
	if err := batch(maxParams/len(schema), individuals, func(individualsInBatch []*api.Individual) (bool, error) {
		args := make([]interface{}, 0)
		b := &strings.Builder{}
		b.WriteString(fmt.Sprintf("INSERT INTO %s", tempTableName))
		b.WriteString(" (")
		for i, col := range columnsOfInterest {
			if i != 0 {
				b.WriteString(",")
			}
			b.WriteString(col)
		}
		b.WriteString(") VALUES ")

		for i, individual := range individualsInBatch {
			if i != 0 {
				b.WriteString(",")
			}
			b.WriteString("(")
			for j, col := range columnsOfInterest {
				if j != 0 {
					b.WriteString(",")
				}
				fieldValue, err := individual.GetFieldValue(col)
				if err != nil {
					return false, err
				}
				if col == constants.DBColumnIndividualID && fieldValue == "" {
					b.WriteString("NULL")
				} else {
					b.WriteString(fmt.Sprintf("$%d", len(args) + 1))
					args = append(args, fieldValue)
				}
			}

			b.WriteString(")")
		}

		var out []*api.Individual
		qry := b.String()

		err := tx.SelectContext(ctx, &out, qry, args...)
		if err != nil {
			return false, err
		}
		return false, err
	}); err != nil {
		return err 
	}
	return nil
}

func buildFileDeduplicationQuery(tempTableName string, columnsOfInterest []string, config deduplication.DeduplicationConfig, schema []DBColumn) string {
	b := &strings.Builder{}

	b.WriteString("SELECT DISTINCT ir.idx idxA, ti.idx idxB")
	b.WriteString((fmt.Sprintf(" FROM %s ir", tempTableName)))
	b.WriteString(fmt.Sprintf(" CROSS JOIN %s ti", tempTableName))
	b.WriteString(" WHERE ir.idx != ti.idx ")

	subQueries := []string{}
	notEmptyPartialChecks := []string{}
	for _, dt := range config.Types {
		if config.Operator == deduplication.LOGICAL_OPERATOR_AND {
			subQueries = append(subQueries, dt.Config.QueryAnd)
			if dt.Config.QueryNotAllEmpty != "" {
				notEmptyPartialChecks = append(notEmptyPartialChecks, dt.Config.QueryNotAllEmpty)
			}
		} else {
			subQueries = append(subQueries, dt.Config.QueryOr)
		}
	}
	if len(subQueries) > 0 {
		b.WriteString(fmt.Sprintf(" AND ((%s))", strings.Join(subQueries, fmt.Sprintf(") %s (", config.Operator))))
	}

	if len(notEmptyPartialChecks) > 0 {
		notAllEmptyQuery := strings.Join(notEmptyPartialChecks, " OR ")
		b.WriteString(" AND ( ")
		b.WriteString(notAllEmptyQuery)
		b.WriteString(" ) ")
	}

	b.WriteString(";")

	s := b.String()
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.Join(strings.Fields(s), " ")

	return s 
}

/*
EXAMPLE:

SELECT DISTINCT  ir.birth_date,ir.email_1,ir.email_2,ir.email_3,ir.first_name,ir.middle_name,ir.last_name,ir.native_name,ir.id

	FROM individual_registrations ir
	CROSS JOIN temp_individuals_<requestId> ti
	WHERE ir.country_id = $1
		AND ir.deleted_at IS NULL
		AND ((ti.birth_date = ir.birth_date)
		AND ((ti.email_1 != '' AND (ti.email_1 = ir.email_1 OR ti.email_1 = ir.email_2 OR ti.email_1 = ir.email_3))
			OR (ti.email_2 != '' AND (ti.email_2 = ir.email_1 OR ti.email_2 = ir.email_2 OR ti.email_2 = ir.email_3))
			OR (ti.email_3 != '' AND (ti.email_3 = ir.email_1 OR ti.email_3 = ir.email_2 OR ti.email_3 = ir.email_3))
			OR (ti.email_1 = '' AND ti.email_2 = '' AND ti.email_3 =''))
		AND (ti.first_name = ir.first_name AND ti.middle_name = ir.middle_name AND ti.last_name = ir.last_name AND ti.native_name = ir.native_name)
		AND (ti.email_1 != '' OR ti.email_2 != '' OR ti.email_3 != '' OR ti.first_name != '' OR ti.middle_name != '' OR ti.last_name != '' OR ti.native_name != '')));
*/
func buildDbDeduplicationQuery(tempTableName string, columnsOfInterest []string, config deduplication.DeduplicationConfig, schema []DBColumn) string {
	b := &strings.Builder{}

	b.WriteString(fmt.Sprintf("SELECT DISTINCT ti.idx, ir.id, ir.%s", constants.DBColumnIndividualLastName))
	for c := range columnsOfInterest {
		if columnsOfInterest[c] == constants.DBColumnIndividualLastName {
			continue
		}
		b.WriteString(fmt.Sprintf(", ir.%s", columnsOfInterest[c]))
	}

	b.WriteString(" FROM individual_registrations ir")
	b.WriteString(fmt.Sprintf(" CROSS JOIN %s ti", tempTableName))
	b.WriteString(" WHERE ir.country_id = $1 AND ir.deleted_at IS NULL")

	subQueries := []string{}
	notEmptyPartialChecks := []string{}
	for _, dt := range config.Types {
		if config.Operator == deduplication.LOGICAL_OPERATOR_AND {
			subQueries = append(subQueries, dt.Config.QueryAnd)
			if dt.Config.QueryNotAllEmpty != "" {
				notEmptyPartialChecks = append(notEmptyPartialChecks, dt.Config.QueryNotAllEmpty)
			}
		} else {
			subQueries = append(subQueries, dt.Config.QueryOr)
		}
	}
	if len(subQueries) > 0 {
		b.WriteString(fmt.Sprintf(" AND ((%s))", strings.Join(subQueries, fmt.Sprintf(") %s (", config.Operator))))
	}

	b.WriteString(" AND (ti.id IS NULL OR ti.id != ir.id)")

	if len(notEmptyPartialChecks) > 0 {
		notAllEmptyQuery := strings.Join(notEmptyPartialChecks, " OR ")
		b.WriteString(" AND ( ")
		b.WriteString(notAllEmptyQuery)
		b.WriteString(" ) ")
	}

	b.WriteString(";")

	s := b.String()
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.Join(strings.Fields(s), " ")

	return s 
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
