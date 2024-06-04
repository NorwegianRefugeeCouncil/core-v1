package db

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"github.com/stretchr/testify/assert"
)

func Seed(ctx context.Context, sqlDb *sqlx.DB) *api.Country {
	countryRepo := NewCountryRepo(sqlDb)
	countryDef := &api.Country{
		Code: "NO",
		Name: "Norway",
		ReadGroup: "read",
		WriteGroup: "write",
	}
	country, err := countryRepo.Put(ctx, countryDef)
	if err != nil {
		log.Fatalf("Failed to seed database: %s", err)
	}
	return country
}

func ClearDatabase(ctx context.Context, sqlDb *sqlx.DB) {
	_, err := sqlDb.ExecContext(ctx, "TRUNCATE TABLE individual_registrations CASCADE")
	if err != nil {
		log.Fatalf("Failed to clear database: %s", err)
	}	
}

func TestDeduplication(t *testing.T) {
	pool, resource := InitTestDocker("5432")
	defer pool.Purge(resource)

	ctx := context.Background()
	sqlDb := OpenDatabaseConnection(ctx, pool, resource, "5432")
	defer sqlDb.Close()

	RunMigrations(ctx, sqlDb)

	country := Seed(ctx, sqlDb)

	tests := []struct {
		name string
		seed []*api.Individual
		individuals []*api.Individual
		deduplicationConfig deduplication.DeduplicationConfig
		wantFile []containers.Set[int]
		wantDb []*api.Individual
	} {
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FullName; [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{},
			individuals: []*api.Individual{
				{
					FullName: "John Doe",
				},
				{
					FullName: "John Doe",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					 deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
			wantFile: []containers.Set[int]{
				containers.NewSet[int](1),
				containers.NewSet[int](0),
			},
			wantDb: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearDatabase(ctx, sqlDb)

			individualRepo := NewIndividualRepo(sqlDb)

			for _, ind := range tt.individuals {
				ind.CountryID = country.ID
			}

			for _, ind := range tt.seed {
				ind.CountryID = country.ID
			}

			individualRepo.PutMany(ctx, tt.seed, constants.IndividualDBColumns)

			fileDupes, dbDupes, err := individualRepo.FindDuplicates(ctx, tt.individuals, tt.deduplicationConfig)
			if err != nil {
				t.Fatalf("Failed to deduplicate: %s", err)
			}

			assert.ElementsMatch(t, tt.wantFile, fileDupes)
			assert.ElementsMatch(t, tt.wantDb, dbDupes)
		})
	}
}