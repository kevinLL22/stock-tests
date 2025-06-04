package db

import (
	"database/sql"
	"fmt"
	"github.com/kevinLL22/stock-tests/migrations"
	"log"

	"github.com/golang-migrate/migrate/v4"
	crdb "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func RunMigrations(databaseURL string) error {
	// 1. source: migrations embedded in the binary
	src, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return fmt.Errorf("iofs: %w", err)
	}

	// 2. Target: driver CockroachDB (sin advisory locks)
	stdDB, err := sql.Open("pgx", databaseURL) // conexión vía pgx stdlib
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	driver, err := crdb.WithInstance(stdDB, &crdb.Config{})
	if err != nil {
		return fmt.Errorf("crdb driver: %w", err)
	}

	// 3. Execute migrations
	m, err := migrate.NewWithInstance("iofs", src, "cockroachdb", driver)
	if err != nil {
		return fmt.Errorf("migrate instance: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("apply: %w", err)
	}

	version, dirty, verr := m.Version()
	if verr != nil {
		log.Printf("Error to obtain version: %v", verr)
	} else if dirty {
		log.Printf("Migrations applid, but in dirty (dirty=true) in version %d", version)
	} else {
		log.Printf("Migrations applied. current version: %d", version)
	}

	return nil
}

func makeStdlibConn(databaseURL string) *sql.DB {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		panic(err) // solo se llama desde runMigrations
	}
	return db
}
