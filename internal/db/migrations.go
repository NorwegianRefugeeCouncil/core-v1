package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"path"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
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
	migrationFromFile("003_countries"),
	migrationFromFile("004_users"),
	migrationFromFile("005_user_defaults"),
	migrationFromFile("006_user_permissions"),
	migrationFromFile("007_global_admin"),
	migrationFromFile("008_individual_country"),
	migrationFromFile("009_remove_permissions"),
	migrationFromFile("010_country_jwt_group"),
	migrationFromFile("011_individual_soft_delete"),
	migrationFromFile("012_new_individual"),
	migrationFromFile("013_individual_changes"),
	migrationFromFile("014_rename_identification_context"),
	migrationFromFile("015_rename_gender"),
	migrationFromFile("016_add_displacement_status_comment_field"),
	migrationFromFile("017_add_office_field"),
	migrationFromFile("018_add_nrc_organisation_to_country"),
	migrationFromFile("019_individual_indices"),
	migrationFromFile("020_rename_nrc_organisation_to_plural"),
	migrationFromFile("021_convert_nrc_organisations_to_array"),
}

// Migrate runs the migrations on the database.
func Migrate(ctx context.Context, db *sqlx.DB) error {

	l := logging.NewLogger(ctx)
	l.Info("migrating database")

	_, err := db.ExecContext(ctx, initScript)
	if err != nil {
		l.Error("failed to initialize database", zap.Error(err))
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
		l.Error("failed to read migrations directory", zap.Error(err))
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		migrationContent, err := migrationFs.ReadFile(path.Join(dir, entry.Name()))
		if err != nil {
			l.Error("failed to read migration file", zap.Error(err))
			panic(err)
		}
		migrationMap[entry.Name()] = string(migrationContent)
	}

	_, err = doInTransaction(ctx, db, func(ctx context.Context, tx *sqlx.Tx) (interface{}, error) {
		for _, m := range migrations {
			l.Info("running migration", zap.String("name", m.name))
			var num int
			err = tx.GetContext(ctx, &num, "SELECT 1 FROM migrations WHERE name = $1", m.name)
			if err == nil {
				l.Info("migration already applied", zap.String("name", m.name))
				continue
			}
			if err != sql.ErrNoRows {
				l.Error("failed to check migration", zap.Error(err))
				return nil, err
			}
			if err := m.up(ctx, tx); err != nil {
				l.Error("failed to run migration", zap.Error(err))
				return nil, err
			}
			if _, err := tx.ExecContext(ctx, "INSERT INTO migrations (name) VALUES ($1)", m.name); err != nil {
				l.Error("failed to insert migration", zap.Error(err))
				return nil, err
			}
		}
		return nil, nil
	})

	if err != nil {
		l.Error("failed to migrate database", zap.Error(err))
		return err
	}

	l.Info("migrated database")

	return nil
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
