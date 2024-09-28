package db

import (
	"database/sql"
	"errors"
	"fmt"

	"goldvault/user-service/internal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to get db instance: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", config.ServiceConfig.Database.Postgres.MigrationsPath),
		config.ServiceConfig.Database.Postgres.Name,
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize db migrations: %w", err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
