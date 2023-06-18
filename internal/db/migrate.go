package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	ErrorCreatingDatabase   = "could not create a database instance"
	ErrorCreatingDriver     = "could not create postgres driver"
	ErrorCreatingMigrations = "could not create migrations"
	SuccessfullyMigratedDB  = "successfully migrated the database"
)

func (d *Database) MigrateDB() error {
	fmt.Println("Migrating database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("%s: %w", ErrorCreatingDriver, err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrorCreatingDatabase, err)
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("%s: %w", ErrorCreatingMigrations, err)
		}
	}
	fmt.Println(SuccessfullyMigratedDB)
	return nil
}
