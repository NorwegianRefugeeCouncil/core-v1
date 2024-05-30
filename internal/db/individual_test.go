package db

import (
	"testing"

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

func TestBuildDeduplicationQuery(t *testing.T) {
	tests := []struct {
		name              string
		tableName         string
		config            deduplication.DeduplicationConfig
		columnsOfInterest []string
		schema            []DBColumn
		wantQuery         string
	}{
		{
			"Deduplication query, no subquery",
			"table_name",
			deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
			[]string{"col_text_1", "col_date"},
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "timestamp", Default: nil},
			},
			"SELECT DISTINCT ir.id, ir.col_text_1,ir.col_date FROM individual_registrations ir CROSS JOIN table_name ti WHERE ir.country_id = $1 AND ir.deleted_at IS NULL AND ((ti.full_name = ir.full_name OR ti.full_name = '')) AND (ti.id IS NULL OR ti.id != ir.id) AND ( ti.full_name != '' ) ;",
		},
		{
			"Deduplication query, id column, deduplicate any, OR subqueries",
			"table_name",
			deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_OR,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
			[]string{"col_text_1", "col_date"},
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "timestamp", Default: nil},
			},
			"SELECT DISTINCT ir.id, ir.col_text_1,ir.col_date FROM individual_registrations ir CROSS JOIN table_name ti WHERE ir.country_id = $1 AND ir.deleted_at IS NULL AND ((ti.full_name != '' AND ti.full_name = ir.full_name) OR ( (ti.identification_number_1 != '' AND (ti.identification_number_1 = ir.identification_number_1 OR ti.identification_number_1 = ir.identification_number_2 OR ti.identification_number_1 = ir.identification_number_3)) OR (ti.identification_number_2 != '' AND (ti.identification_number_2 = ir.identification_number_1 OR ti.identification_number_2 = ir.identification_number_2 OR ti.identification_number_2 = ir.identification_number_3)) OR (ti.identification_number_3 != '' AND (ti.identification_number_3 = ir.identification_number_1 OR ti.identification_number_3 = ir.identification_number_2 OR ti.identification_number_3 = ir.identification_number_3)) )) AND (ti.id IS NULL OR ti.id != ir.id);",
		},
		{
			"Deduplication query, id column, deduplicate all, OR + AND subqueries",
			"table_name",
			deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameEmails],
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
				},
			},
			[]string{},
			[]DBColumn{
				{Name: "id", SQLType: "uuid", Default: nil},
				{Name: "col_text_1", SQLType: "text", Default: pointers.String("")},
				{Name: "col_text_2", SQLType: "text", Default: nil},
				{Name: "col_date", SQLType: "timestamp", Default: nil},
			},
			"SELECT DISTINCT ir.id, FROM individual_registrations ir CROSS JOIN table_name ti WHERE ir.country_id = $1 AND ir.deleted_at IS NULL AND (( (ti.email_1 != '' AND (ti.email_1 = ir.email_1 OR ti.email_1 = ir.email_2 OR ti.email_1 = ir.email_3)) OR (ti.email_2 != '' AND (ti.email_2 = ir.email_1 OR ti.email_2 = ir.email_2 OR ti.email_2 = ir.email_3)) OR (ti.email_3 != '' AND (ti.email_3 = ir.email_1 OR ti.email_3 = ir.email_2 OR ti.email_3 = ir.email_3)) OR (ti.email_1 = '' AND ti.email_2 = '' AND ti.email_3 ='') ) AND ( (ti.first_name = ir.first_name AND ti.middle_name = ir.middle_name AND ti.last_name = ir.last_name AND ti.native_name = ir.native_name) OR (ti.first_name = '' AND ti.middle_name = '' AND ti.last_name = '' AND ti.native_name = '') )) AND (ti.id IS NULL OR ti.id != ir.id) AND ( ti.email_1 != '' OR ti.email_2 != '' OR ti.email_3 != '' OR ti.first_name != '' OR ti.middle_name != '' OR ti.last_name != '' OR ti.native_name != '' ) ;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := buildDbDeduplicationQuery(tt.tableName, tt.columnsOfInterest, tt.config, tt.schema)
			assert.Equal(t, tt.wantQuery, query)
		})
	}
}
