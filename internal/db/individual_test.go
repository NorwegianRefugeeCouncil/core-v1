package db

import (
	"testing"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/nrc-no/notcore/internal/utils/pointers"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"github.com/stretchr/testify/assert"
)

func TestBuildTableSchemaQuery(t *testing.T) {
	t.Run("Get Table schema query", func(t *testing.T) {
		query := buildTableSchemaQuery()
		assert.Equal(t, query, "SELECT column_name as Name, udt_name as SQLType, column_default as Default FROM information_schema.columns WHERE table_name = 'individual_registrations';")
	})
}

func TestBuildCreateTempTableQuery(t *testing.T) {
	tests := []struct {
		name              string
		tableName         string
		schema            []DBColumn
		columnsOfInterest []string
		wantQuery         string
	}{
		{
			"Create temp table query",
			"table_name",
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "timestamp", Default: nil},
			},
			[]string{"col_text_1", "col_date"},
			"CREATE TEMPORARY TABLE table_name (col_text_1 text,col_date timestamp) ON COMMIT DROP;",
		},
		{
			"Create temp table query",
			"table_name",
			[]DBColumn{},
			[]string{"col_text_1", "col_date"},
			"CREATE TEMPORARY TABLE table_name () ON COMMIT DROP;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := buildCreateTempTableQuery(tt.tableName, tt.schema, tt.columnsOfInterest)
			assert.Equal(t, tt.wantQuery, query)
		})
	}
}

func TestBuildInsertIndividualsQuery(t *testing.T) {
	tests := []struct {
		name              string
		tableName         string
		schema            []DBColumn
		df                dataframe.DataFrame
		columnsOfInterest []string
		uploadHasIdColumn bool
		wantQuery         string
		wantArgs          []interface{}
	}{
		{
			"Insert individuals query, skip irrelevant columns",
			"table_name",
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "date", Default: nil},
			},
			dataframe.New(
				series.New([]string{"c", "b", "a"}, series.String, "col_text_1"),
				series.New([]string{"2", "1", ""}, series.String, "col_text_2"),
				series.New([]string{"", "0000", "2003-12-30"}, series.String, "col_date"),
			),
			[]string{"col_text_1", "col_date"},
			false,
			"INSERT INTO table_name SELECT * FROM UNNEST($1::text[],$2::date[]);",
			[]interface{}{
				pq.Array([]string{"c", "b", "a"}),
				pq.Array([]*time.Time{nil, nil, pointers.Time(time.Date(2003, 12, 30, 0, 0, 0, 0, time.UTC))}),
			},
		},
		{
			"Insert individuals query, add id column, skip non-existent columns",
			"table_name",
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "date", Default: nil},
			},
			dataframe.New(
				series.New([]string{"01234567-0123-4567-8901-012345678901", "01234567-0123-4567-8901-012345678902"}, series.String, "id"),
				series.New([]string{"b", "a"}, series.String, "col_text_1"),
				series.New([]string{"1", ""}, series.String, "col_text_2"),
				series.New([]string{"", "3"}, series.String, "col_date"),
			),
			[]string{"other", "no"},
			true,
			"INSERT INTO table_name SELECT * FROM UNNEST($1::uuid[]);",
			[]interface{}{pq.Array([]uuid.UUID{uuid.MustParse("01234567-0123-4567-8901-012345678901"), uuid.MustParse("01234567-0123-4567-8901-012345678902")})},
		},
		{
			"Insert individuals query, add id column, fill with empty column if no data exists",
			"table_name",
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "date", Default: nil},
			},
			dataframe.New(
				series.New([]string{"01234567-0123-4567-8901-012345678901", "01234567-0123-4567-8901-012345678902"}, series.String, "id"),
				series.New([]string{"b", "a"}, series.String, "col"),
				series.New([]string{"1", ""}, series.String, "col_text_2"),
				series.New([]string{"", "3"}, series.String, "col_date"),
			),
			[]string{"col_text_1", "no"},
			true,
			"INSERT INTO table_name SELECT * FROM UNNEST($1::uuid[],$2::text[]);",
			[]interface{}{
				pq.Array([]uuid.UUID{uuid.MustParse("01234567-0123-4567-8901-012345678901"), uuid.MustParse("01234567-0123-4567-8901-012345678902")}),
				pq.Array(nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, args := buildInsertIndividualsQuery(tt.tableName, tt.schema, tt.df, tt.columnsOfInterest, tt.uploadHasIdColumn)
			assert.Equal(t, tt.wantArgs, args)
			assert.Equal(t, tt.wantQuery, query)
		})
	}
}

func TestBuildDeduplicationQuery(t *testing.T) {
	tests := []struct {
		name              string
		tableName         string
		config            deduplication.DeduplicationConfig
		columnsOfInterest []string
		uploadHasIdColumn bool
		schema            []DBColumn
		wantQuery         string
	}{
		{
			"Deduplication query, no id column, OR subquery",
			"table_name",
			deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
			[]string{"col_text_1", "col_date"},
			false,
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "timestamp", Default: nil},
			},
			"SELECT DISTINCT ir.col_text_1,ir.col_date,ir.id FROM individual_registrations ir CROSS JOIN table_name ti WHERE ir.country_id = $1 AND ir.deleted_at IS NULL AND (ti.full_name = ir.full_name);",
		},
		{
			"Deduplication query, id column, deduplicate any, OR subqueries",
			"table_name",
			deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_OR,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
			[]string{"col_text_1", "col_date"},
			true,
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "timestamp", Default: nil},
			},
			"SELECT DISTINCT ir.col_text_1,ir.col_date,ir.id FROM individual_registrations ir CROSS JOIN table_name ti WHERE ir.country_id = $1 AND ir.deleted_at IS NULL AND (ti.full_name = ir.full_name) OR ((ti.identification_number_1 != '' AND ti.identification_number_1 = ir.identification_number_1) OR (ti.identification_number_2 != '' AND ti.identification_number_2 = ir.identification_number_2) OR (ti.identification_number_3 != '' AND ti.identification_number_3 = ir.identification_number_3)) AND ti.id::uuid NOT IN (SELECT id FROM individual_registrations);",
		},
		{
			"Deduplication query, id column, deduplicate all, OR + AND subqueries",
			"table_name",
			deduplication.DeduplicationConfig{
				deduplication.LOGICAL_OPERATOR_AND,
				[]deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameEmails],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
				},
			},
			[]string{},
			true,
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "timestamp", Default: nil},
			},
			"SELECT DISTINCT ir.id FROM individual_registrations ir CROSS JOIN table_name ti WHERE ir.country_id = $1 AND ir.deleted_at IS NULL AND ((ti.email_1 != '' AND ti.email_1 = ir.email_1) OR (ti.email_2 != '' AND ti.email_2 = ir.email_2) OR (ti.email_3 != '' AND ti.email_3 = ir.email_3)) AND (ti.first_name = ir.first_name AND ti.middle_name = ir.middle_name AND ti.last_name = ir.last_name AND ti.native_name = ir.native_name AND (ti.first_name != '' OR ti.middle_name != '' OR ti.last_name != '' OR ti.native_name != '')) AND ti.id::uuid NOT IN (SELECT id FROM individual_registrations);",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := buildDeduplicationQuery(tt.tableName, tt.columnsOfInterest, tt.config, tt.uploadHasIdColumn, tt.schema)
			assert.Equal(t, tt.wantQuery, query)
		})
	}
}
