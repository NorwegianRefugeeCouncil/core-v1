package db

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
)

func TestBatch(t *testing.T) {
	tests := []struct {
		name      string
		batchSize int
		all       []string
		want      [][]string
		wantErr   bool
	}{
		{
			name:      "empty",
			batchSize: 10,
			all:       []string{},
			want:      [][]string{},
		}, {
			name:      "batch size 1",
			batchSize: 1,
			all:       []string{"a", "b", "c"},
			want:      [][]string{{"a"}, {"b"}, {"c"}},
		}, {
			name:      "batch size 2",
			batchSize: 2,
			all:       []string{"a", "b", "c"},
			want:      [][]string{{"a", "b"}, {"c"}},
		}, {
			name:      "batch size 3",
			batchSize: 3,
			all:       []string{"a", "b", "c"},
			want:      [][]string{{"a", "b", "c"}},
		}, {
			name:      "batch size 4",
			batchSize: 4,
			all:       []string{"a", "b", "c"},
			want:      [][]string{{"a", "b", "c"}},
		}, {
			name:      "batch size 0",
			batchSize: 0,
			all:       []string{"a", "b", "c"},
			wantErr:   true,
		}, {
			name:      "batch size -1",
			batchSize: -1,
			all:       []string{"a", "b", "c"},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = make([][]string, 0)
			err := batch(tt.batchSize, tt.all, func(batch []string) (bool, error) {
				got = append(got, batch)
				return false, nil
			})

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

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