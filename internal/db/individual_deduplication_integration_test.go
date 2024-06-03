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
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
)

func InitTestDocker(exposedPort string) (*dockertest.Pool, *dockertest.Resource) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14",
		Env: []string{
			"listen_addresses = '*'",
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=core",
		},
		ExposedPorts: []string{exposedPort},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {
				{HostIP: "0.0.0.0", HostPort: exposedPort},
			},
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}// Important option when container crash and you want to debug container
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := resource.Expire(30); err != nil { // Tell docker to hard kill the container in 30 seconds
		fmt.Println("Could not set container expiration", err)
	}

	// retry if container is not ready
	pool.MaxWait = 30 * time.Second
	if err = pool.Retry(func() error {
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return pool, resource
}

func OpenDatabaseConnection(ctx context.Context, pool *dockertest.Pool, resource *dockertest.Resource, exposedPort string) *sqlx.DB {
	databaseDriver := "postgres"
	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseDSN := fmt.Sprintf("postgres://postgres:postgres@%s/core?sslmode=disable", hostAndPort)

	retries := 5
	sqlDb, err := sqlx.ConnectContext(ctx, databaseDriver, databaseDSN)
	
  // Sometimes it happens that after first time container is not ready. 
  // It's always better to create retry if necessary and be sure that tests run without problems
  for err != nil {
		if retries > 1 {
			retries--
			time.Sleep(1 * time.Second)
			sqlDb, err = sqlx.ConnectContext(ctx, databaseDriver, databaseDSN)
			continue
		}

		if err := pool.Purge(resource); err != nil {
			fmt.Printf("Could not purge resource: %s", err)
		}

		log.Panic("Fatal error in connection: ", err, resource.GetBoundIP("5432/tcp"))
	}

	return sqlDb 
}

func RunMigrations(ctx context.Context, sqlDb *sqlx.DB) {
	if err := Migrate(ctx, sqlDb); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}
}

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