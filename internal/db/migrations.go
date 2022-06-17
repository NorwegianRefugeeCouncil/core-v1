package db

import (
	"context"
	"embed"
	"fmt"
	"path"

	"github.com/jmoiron/sqlx"
)

// migration is a single migration.
type migration struct {
	// name is the name of the migration.
	name string
	// up is the SQL to run for the up migration.
	up func(ctx context.Context, tx *sqlx.Tx) error
	// down is the SQL to run for the down migration.
	down func(ctx context.Context, tx *sqlx.Tx) error
}

// migrations contains the list of migrations to run.
var migrations = []migration{
	migrationFromFile("initial"),
}

// Migrate runs the migrations on the database.
func Migrate(ctx context.Context, db *sqlx.DB) error {

	var driver = ""
	if db.DriverName() == "sqlite3" {
		driver = "sqlite"
	} else if db.DriverName() == "postgres" {
		driver = "postgres"
	} else {
		return fmt.Errorf("unsupported driver: %s", db.DriverName())
	}

	dir := fmt.Sprintf("migrations/%s", driver)
	entries, err := migrationFs.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		migrationContent, err := migrationFs.ReadFile(path.Join(dir, entry.Name()))
		if err != nil {
			panic(err)
		}
		migrationMap[entry.Name()] = string(migrationContent)
	}

	_, err = doInTransaction(ctx, db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		for _, m := range migrations {
			if err := m.up(ctx, tx); err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	return err
}

//go:embed migrations/**/*.sql
var migrationFs embed.FS

// migrationMap contains a map of migration filenames to SQL statements.
var migrationMap = map[string]string{}

// migrationFromFile creates a migration from the migrationFs
// name correspond to the migration file name <name>.up.sql or <name>.down.sql
func migrationFromFile(name string) migration {
	upName := name + ".up.sql"
	downName := name + ".down.sql"
	return migration{
		name: name,
		up: func(ctx context.Context, tx *sqlx.Tx) error {
			_, err := tx.ExecContext(ctx, migrationMap[upName])
			return err
		},
		down: func(ctx context.Context, tx *sqlx.Tx) error {
			_, err := tx.ExecContext(ctx, migrationMap[downName])
			return err
		},
	}
}
