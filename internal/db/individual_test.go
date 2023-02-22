package db

import (
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
		driver               string
		want                 string
	}{
		{
			name:                "3 existing individuals, 1 unchecked individual, 1 deduplication type",
			existingIndividuals: containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:  []DeduplicationOptionName{DeduplicationOptionNameIds},
			uncheckedIndividuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "1", IdentificationNumber2: "2", IdentificationNumber3: "3"},
			},
			driver: "postgres",
			want:   "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ('{\"1\",\"2\",\"3\"}'::uuid[])) AND deleted_at IS NULL AND (identification_number_1 IN ('1') OR identification_number_2 IN ('2') OR identification_number_3 IN ('3'))",
		},
		{
			name:                "3 existing individuals, 2 unchecked individuals, 1 deduplication type",
			existingIndividuals: containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:  []DeduplicationOptionName{DeduplicationOptionNameIds},
			uncheckedIndividuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "1", IdentificationNumber2: "2", IdentificationNumber3: "3"},
				{ID: "4", IdentificationNumber1: "4", IdentificationNumber2: "5", IdentificationNumber3: "6"},
			},
			driver: "postgres",
			want:   "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ('{\"1\",\"2\",\"3\"}'::uuid[])) AND deleted_at IS NULL AND (identification_number_1 IN ('1','4') OR identification_number_2 IN ('2','5') OR identification_number_3 IN ('3','6'))",
		},
		{
			name:                "3 existing individuals, 2 unchecked individuals, 2 deduplication types, empty values",
			existingIndividuals: containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:  []DeduplicationOptionName{DeduplicationOptionNameIds, DeduplicationOptionNameFullName},
			uncheckedIndividuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "1", IdentificationNumber2: "2", IdentificationNumber3: "3"},
				{ID: "4", IdentificationNumber1: "4", IdentificationNumber2: "5", IdentificationNumber3: "6"},
			},
			driver: "postgres",
			want:   "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ('{\"1\",\"2\",\"3\"}'::uuid[])) AND deleted_at IS NULL AND (identification_number_1 IN ('1','4') OR identification_number_2 IN ('2','5') OR identification_number_3 IN ('3','6'))",
		},
		{
			name:                "3 existing individuals, 2 unchecked individuals, 2 deduplication types",
			existingIndividuals: containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:  []DeduplicationOptionName{DeduplicationOptionNameIds, DeduplicationOptionNameFullName},
			uncheckedIndividuals: []*api.Individual{
				{ID: "1", IdentificationNumber1: "1", IdentificationNumber2: "2", IdentificationNumber3: "3", FullName: "John Doe"},
				{ID: "4", IdentificationNumber1: "4", IdentificationNumber2: "5", IdentificationNumber3: "6", FullName: "Jane Doe"},
			},
			driver: "postgres",
			want:   "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ('{\"1\",\"2\",\"3\"}'::uuid[])) AND deleted_at IS NULL AND (identification_number_1 IN ('1','4') OR identification_number_2 IN ('2','5') OR identification_number_3 IN ('3','6')) AND (full_name IN ('John Doe','Jane Doe'))",
		},
		{
			name:                 "3 existing individuals, 0 unchecked individuals, 1 deduplication type",
			existingIndividuals:  containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:   []DeduplicationOptionName{DeduplicationOptionNameIds},
			uncheckedIndividuals: []*api.Individual{},
			driver:               "postgres",
			want:                 "SELECT * FROM individual_registrations WHERE id NOT IN (SELECT * FROM UNNEST ('{\"1\",\"2\",\"3\"}'::uuid[])) AND deleted_at IS NULL",
		},
		{
			name:                 "sqlite3",
			existingIndividuals:  containers.NewStringSet("1", "2", "3"),
			deduplicationTypes:   []DeduplicationOptionName{DeduplicationOptionNameIds},
			uncheckedIndividuals: []*api.Individual{},
			driver:               "sqlite3",
			want:                 "SELECT * FROM individual_registrations WHERE id NOT IN ('1','2','3') AND deleted_at IS NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := buildDeduplicationQuery(tt.driver, tt.existingIndividuals, tt.uncheckedIndividuals, tt.deduplicationTypes)
			assert.Equal(t, tt.want, query)
		})
	}
}
