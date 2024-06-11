package db

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) Migrate() error {
	fmt.Println("Migrating database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println("failed to create migrate instance: %w", err)
		return err
	}
	content, err := fs.ReadFile(os.DirFS("migrations"), "0002_create_comments_table.up.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}
	fmt.Printf("migrations path: %s\n", m)
	fmt.Printf("migrations path: %s\n", content)
	m.Force(1)
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("failed to migrate database ErrNoChange: %w", err)
		}
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	fmt.Println("Successfully migrated database")

	return nil
}
