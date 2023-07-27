package db

import (
	"github.com/lib/pq"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var individuals = []*api.Individual{
	{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3", FirstName: "FN1", MiddleName: "", LastName: "LN1", NativeName: "NN1", Email1: "123"},
	{ID: "2", IdentificationNumber1: "ID4", IdentificationNumber2: "ID5", IdentificationNumber3: "ID6", FirstName: "FN2", MiddleName: "MN2", LastName: "LN2", NativeName: "NN2", Email2: "456"},
	{ID: "3", IdentificationNumber1: "ID7", IdentificationNumber2: "ID8", FirstName: "FN3", LastName: "LN3", MiddleName: ""},
}

func TestGetEmptyValuesQuery(t *testing.T) {
	tests := []struct {
		name               string
		deduplicationTypes []deduplication.DeduplicationTypeName
		wantValues         string
	}{
		{
			name:               "type: Names, IDs, Emails",
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameEmails},
			wantValues:         "first_name = '' AND middle_name = '' AND last_name = '' AND native_name = '' AND identification_number_1 = '' AND identification_number_2 = '' AND identification_number_3 = '' AND email_1 = '' AND email_2 = '' AND email_3 = ''",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := getEmptyValuesQuery(tt.deduplicationTypes)
			assert.Equal(t, tt.wantValues, params)
		})
	}
}

func TestCollectParams(t *testing.T) {
	tests := []struct {
		name               string
		deduplicationTypes []deduplication.DeduplicationTypeName
		wantValues         QueryArgs
	}{
		{
			name:               "type: Names, IDs",
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameEmails, deduplication.DeduplicationTypeNameFullName},
			wantValues: QueryArgs{
				And: AndTypeArgsGroups{
					"Names": individuals,
				},
				Or: OrTypeArgsGroups{
					"Ids": ColumnArgsGroups{
						"identification_number_1": {"ID1", "ID4", "ID7", "ID2", "ID5", "ID8", "ID3", "ID6"},
						"identification_number_2": {"ID1", "ID4", "ID7", "ID2", "ID5", "ID8", "ID3", "ID6"},
						"identification_number_3": {"ID1", "ID4", "ID7", "ID2", "ID5", "ID8", "ID3", "ID6"},
					},
					"Emails": ColumnArgsGroups{
						"email_1": {"123", "456"},
						"email_2": {"123", "456"},
						"email_3": {"123", "456"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := collectArgs(individuals, tt.deduplicationTypes)
			assert.Equal(t, tt.wantValues, params)
		})
	}
}

func TestFillOrQueryWithParameters(t *testing.T) {
	tests := []struct {
		name             string
		queryBuilder     *strings.Builder
		values           ColumnArgsGroups
		wantArgs         []interface{}
		wantedQueryParts containers.Set[string]
	}{
		{
			name:         "type: IDs",
			queryBuilder: &strings.Builder{},
			values: ColumnArgsGroups{
				"identification_number_1": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
				"identification_number_2": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
				"identification_number_3": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
			},
			wantArgs: []interface{}{
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
			},
			wantedQueryParts: containers.NewSet[string]([]string{
				"identification_number_1 IN (SELECT * FROM UNNEST ($1::text[]))",
				"identification_number_2 IN (SELECT * FROM UNNEST ($1::text[]))",
				"identification_number_3 IN (SELECT * FROM UNNEST ($1::text[]))",
			}...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []interface{}{}
			query := ""
			query, args = getOrSubQueriesWithArgs(args, tt.values)
			assert.Equal(t, tt.wantArgs, args)

			assert.Equal(t, containers.NewSet[string](strings.Split(query, " OR ")...), tt.wantedQueryParts)
		})
	}
}

func TestFillAndQueryWithParameters(t *testing.T) {
	tests := []struct {
		name             string
		queryBuilder     *strings.Builder
		values           RowArgsGroups
		wantArgs         []interface{}
		wantedQueryParts [][]string
		typeKey          deduplication.DeduplicationTypeName
	}{
		{
			name:         "type: Names",
			queryBuilder: &strings.Builder{},
			values:       individuals,
			typeKey:      deduplication.DeduplicationTypeNameNames,
			wantArgs: []interface{}{
				"FN1", "LN1", "NN1", "FN2", "MN2", "LN2", "NN2", "FN3", "LN3",
			},
			wantedQueryParts: [][]string{
				{"first_name = $1", "last_name = $2", "native_name = $3"},
				{"first_name = $4", "middle_name = $5", "last_name = $6", "native_name = $7"},
				{"first_name = $8", "last_name = $9"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []interface{}{}
			query := ""
			query, args = getAndSubQueriesWithArgs(args, tt.values, deduplication.DeduplicationTypeNameNames)
			assert.Equal(t, tt.wantArgs, args)

			query = strings.TrimPrefix(query, "(")
			query = strings.TrimSuffix(query, ")")
			queries := strings.Split(query, ") OR (")
			for q := 0; q < len(queries); q++ {
				assert.Equal(t, strings.Split(queries[q], " AND "), tt.wantedQueryParts[q])
			}
		})
	}
}
