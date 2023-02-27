package db

import (
	"github.com/lib/pq"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildDeduplicationQuery(t *testing.T) {
	tests := []struct {
		name                 string
		existingIndividuals  containers.StringSet
		deduplicationTypes   []DeduplicationOptionName
		uncheckedIndividuals []*api.Individual
		wantQuery            string
		wantArgs             []interface{}
	}{
		{
			name:                "3 existing individuals, 1 unchecked individual, 1 deduplication type",
			existingIndividuals: containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:  []DeduplicationOptionName{DeduplicationOptionNameIds},
			uncheckedIndividuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3"},
			},
			wantQuery: "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ($1::uuid[])) AND deleted_at IS NULL AND (identification_number_1 IN (SELECT * FROM UNNEST ($2::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($3::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($4::text[])))",
			wantArgs: []interface{}{
				pq.Array([]string{"1", "2", "3"}),
				pq.Array([]string{"ID1", "ID2", "ID3"}),
				pq.Array([]string{"ID1", "ID2", "ID3"}),
				pq.Array([]string{"ID1", "ID2", "ID3"}),
			},
		},
		{
			name:                "3 existing individuals, 2 unchecked individuals, 1 deduplication type",
			existingIndividuals: containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:  []DeduplicationOptionName{DeduplicationOptionNameIds},
			uncheckedIndividuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3"},
				{ID: "4", IdentificationNumber1: "ID4", IdentificationNumber2: "ID5", IdentificationNumber3: "ID6"},
			},
			wantQuery: "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ($1::uuid[])) AND deleted_at IS NULL AND (identification_number_1 IN (SELECT * FROM UNNEST ($2::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($3::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($4::text[])))",
			wantArgs: []interface{}{
				pq.Array([]string{"1", "2", "3"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
			},
		},
		{
			name:                "3 existing individuals, 2 unchecked individuals, 2 deduplication types, empty values",
			existingIndividuals: containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:  []DeduplicationOptionName{DeduplicationOptionNameIds, DeduplicationOptionNameFullName},
			uncheckedIndividuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3"},
				{ID: "4", IdentificationNumber1: "ID4", IdentificationNumber2: "ID5", IdentificationNumber3: "ID6"},
			},
			wantQuery: "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ($1::uuid[])) AND deleted_at IS NULL AND (identification_number_1 IN (SELECT * FROM UNNEST ($2::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($3::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($4::text[])))",
			wantArgs: []interface{}{
				pq.Array([]string{"1", "2", "3"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
				pq.Array([]string{"ID1", "ID4", "ID2", "ID5", "ID3", "ID6"}),
			},
		},
		{
			name:                "3 existing individuals, 2 unchecked individuals, 2 deduplication types",
			existingIndividuals: containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:  []DeduplicationOptionName{DeduplicationOptionNameIds, DeduplicationOptionNameNames},
			uncheckedIndividuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "ID1", IdentificationNumber2: "ID2", IdentificationNumber3: "ID3", FirstName: "John", LastName: "Doe"},
				{ID: "4", IdentificationNumber1: "ID4", IdentificationNumber2: "ID5", IdentificationNumber3: "ID6", FirstName: "Jane", LastName: "Doe"},
			},
			wantQuery: "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ($1::uuid[])) AND deleted_at IS NULL AND (identification_number_1 IN (SELECT * FROM UNNEST ($2::text[])) OR identification_number_2 IN (SELECT * FROM UNNEST ($3::text[])) OR identification_number_3 IN (SELECT * FROM UNNEST ($4::text[]))) AND (first_name IN (SELECT * FROM UNNEST ($5::text[])) AND last_name IN (SELECT * FROM UNNEST ($6::text[])))",
			wantArgs: []interface{}{
				pq.Array([]string{"1", "2", "3"}),
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
			query, args := buildDeduplicationQuery(tt.existingIndividuals, tt.uncheckedIndividuals, tt.deduplicationTypes)
			assert.Equal(t, tt.wantQuery, query)
			assert.Equal(t, tt.wantArgs, args)
		})
	}
}
