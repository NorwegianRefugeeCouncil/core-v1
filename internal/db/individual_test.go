package db

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var individuals = []*api.Individual{
	{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3", FirstName: "FN1", MiddleName: "MN1", LastName: "LN1", NativeName: "NN1", FullName: "FN1 MN1 LN1"},
	{ID: "2", IdentificationNumber1: "ID4", IdentificationNumber2: "ID5", IdentificationNumber3: "ID6", FirstName: "FN2", MiddleName: "MN2", LastName: "LN2", NativeName: "NN2", FullName: "FN2 MN2 LN2"},
}

func TestCollectParamsForOrQuery(t *testing.T) {
	tests := []struct {
		name               string
		deduplicationTypes deduplication.DeduplicationTypeValue
		wantValues         columnValues
	}{
		{
			name:               "type: IDs",
			deduplicationTypes: deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds].Config,
			wantValues: columnValues{
				"identification_number_1": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
				"identification_number_2": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
				"identification_number_3": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
			},
		},
		{
			name:               "type: Full Name",
			deduplicationTypes: deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName].Config,
			wantValues: columnValues{
				"full_name": {"FN1 MN1 LN1", "FN2 MN2 LN2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := map[string][]string{}
			values = collectParamsForOrQuery(individuals, tt.deduplicationTypes, values)
			assert.Equal(t, tt.wantValues, values)
		})
	}
}

func TestCollectParamsForAndQuery(t *testing.T) {
	tests := []struct {
		name               string
		deduplicationTypes deduplication.DeduplicationTypeValue
		wantValues         columnValues
	}{
		{
			name:               "type: Names",
			deduplicationTypes: deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames].Config,
			wantValues: columnValues{
				"first_name":  {"FN1", "FN2"},
				"last_name":   {"LN1", "LN2"},
				"middle_name": {"MN1", "MN2"},
				"native_name": {"NN1", "NN2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := map[string][]string{}
			values = collectParamsForAndQuery(individuals, tt.deduplicationTypes, values)
			assert.Equal(t, tt.wantValues, values)
		})
	}
}

func TestCollectParams(t *testing.T) {
	tests := []struct {
		name               string
		deduplicationTypes []deduplication.DeduplicationTypeName
		wantValues         queryValues
	}{
		{
			name:               "type: Names, IDs",
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds},
			wantValues: queryValues{
				"Names": columnValues{
					"first_name":  {"FN1", "FN2"},
					"last_name":   {"LN1", "LN2"},
					"middle_name": {"MN1", "MN2"},
					"native_name": {"NN1", "NN2"},
				},
				"Ids": columnValues{
					"identification_number_1": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
					"identification_number_2": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
					"identification_number_3": {"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := collectParams(individuals, tt.deduplicationTypes)
			assert.Equal(t, tt.wantValues, params)
		})
	}
}

func TestFillQueryWithParameters(t *testing.T) {
	a := pq.Array([]string{"FN1", "FN2"})
	b, c := a.Value()
	fmt.Println(b, c)
	tests := []struct {
		name         string
		queryBuilder *strings.Builder
		params       queryValues
		wantArgs     []interface{}
		wantQuery    string
	}{
		{
			name:         "type: Names, IDs",
			queryBuilder: &strings.Builder{},
			params:       collectParams(individuals, []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameNames, deduplication.DeduplicationTypeNameIds}),
			wantArgs: []interface{}{
				pq.Array([]string{"LN1", "LN2"}),
				pq.Array([]string{"FN1", "FN2"}),
				pq.Array([]string{"LN1", "LN2"}),
				pq.Array([]string{"MN1", "MN2"}),
				pq.Array([]string{"NN1", "NN2"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
			},
			wantQuery: " AND (last_name IN (SELECT * FROM UNNEST ($1::text[])) AND native_name IN (SELECT * FROM UNNEST ($2::text[])) AND first_name IN (SELECT * FROM UNNEST ($3::text[])) AND middle_name IN (SELECT * FROM UNNEST ($4::text[]))) AND (identification_number_1 IN (SELECT * FROM UNNEST ($5::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($6::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($7::text[])))\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := fillQueryWithParameters(tt.params, tt.queryBuilder)
			assert.Equal(t, tt.wantArgs, args)
			assert.Equal(t, tt.wantQuery, tt.queryBuilder.String())
		})
	}
}

func TestBuildDeduplicationQuery(t *testing.T) {
	tests := []struct {
		name               string
		deduplicationTypes []deduplication.DeduplicationTypeName
		individuals        []*api.Individual
		wantQuery          string
		wantArgs           []interface{}
	}{
		{
			name:               "3 existing individuals, 1 unchecked individual, 1 deduplication type",
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds},
			individuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3"},
			},
			wantQuery: "SELECT * FROM individual_registrations WHERE country_id = 'countryId' AND deleted_at IS NULL AND (identification_number_1 IN (SELECT * FROM UNNEST ($1::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($2::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($3::text[])))",
			wantArgs: []interface{}{
				pq.Array([]string{"ID1", "ID2", "ID3"}),
				pq.Array([]string{"ID1", "ID2", "ID3"}),
				pq.Array([]string{"ID1", "ID2", "ID3"}),
			},
		},
		{
			name:               "3 existing individuals, 2 unchecked individuals, 1 deduplication type",
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds},
			individuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3"},
				{ID: "4", IdentificationNumber1: "ID4", IdentificationNumber2: "ID5", IdentificationNumber3: "ID6"},
			},
			wantQuery: "SELECT * FROM individual_registrations WHERE country_id = 'countryId' AND deleted_at IS NULL AND (identification_number_1 IN (SELECT * FROM UNNEST ($1::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($2::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($3::text[])))",
			wantArgs: []interface{}{
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
			},
		},
		{
			name:               "3 existing individuals, 2 unchecked individuals, 2 deduplication types, empty values",
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameFullName},
			individuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3"},
				{ID: "4", IdentificationNumber1: "ID4", IdentificationNumber2: "ID5", IdentificationNumber3: "ID6"},
			},
			wantQuery: "SELECT * FROM individual_registrations WHERE country_id = 'countryId' AND deleted_at IS NULL AND (identification_number_1 IN (SELECT * FROM UNNEST ($1::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($2::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($3::text[])))",
			wantArgs: []interface{}{
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
			},
		},
		{
			name:               "3 existing individuals, 2 unchecked individuals, 2 deduplication types",
			deduplicationTypes: []deduplication.DeduplicationTypeName{deduplication.DeduplicationTypeNameIds, deduplication.DeduplicationTypeNameNames},
			individuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3", FirstName: "John", LastName: "Doe"},
				{ID: "4", IdentificationNumber1: "ID4", IdentificationNumber2: "ID5", IdentificationNumber3: "ID6", FirstName: "Jane", LastName: "Doe"},
			},
			wantQuery: "SELECT * FROM individual_registrations WHERE country_id = 'countryId' AND deleted_at IS NULL AND (identification_number_1 IN (SELECT * FROM UNNEST ($1::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($2::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($3::text[]))) AND (first_name IN (SELECT * FROM UNNEST ($4::text[])) AND last_name IN (SELECT * FROM UNNEST ($5::text[])))",
			wantArgs: []interface{}{
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"John", "Jane"}),
				pq.Array([]string{"Doe", "Doe"}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, args := buildDeduplicationQuery("countryId", tt.individuals, tt.deduplicationTypes)
			assert.Equal(t, tt.wantQuery, query)
			assert.Equal(t, tt.wantArgs, args)
		})
	}
}
