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
}

// migrations contains the list of migrations to run.
var migrations = []migration{
	migrationFromFile("001_initial"),
	migrationFromFile("002_global_intake"),
}

// Migrate runs the migrations on the database.
func Migrate(ctx context.Context, db *sqlx.DB) error {

	_, err := db.ExecContext(ctx, initScript)
	if err != nil {
		return err
	}

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
			var num int
			err = tx.GetContext(ctx, &num, "SELECT 1 FROM migrations WHERE name = $1", m.name)
			if err != nil {
				return nil, err
			}
			if num > 0 {
				continue
			}
			if err := m.up(ctx, tx); err != nil {
				return nil, err
			}
			if _, err := tx.ExecContext(ctx, "INSERT INTO migrations (name) VALUES ($1)", m.name); err != nil {
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
// name correspond to the migration file name <name>.up.sql
func migrationFromFile(name string) migration {
	upName := name + ".up.sql"
	return migration{
		name: name,
		up: func(ctx context.Context, tx *sqlx.Tx) error {
			if migrationMap[upName] == "" {
				return fmt.Errorf("migration %s not found", upName)
			}
			_, err := tx.ExecContext(ctx, migrationMap[upName])
			return err
		},
	}
}

var initScript = `
CREATE TABLE IF NOT EXISTS migrations
(
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY (name)
)
`
