package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// flags -c=connection_string -m=path_to_migrations
const (
	connection_string  = "postgres://postgres:230704@localhost:5432/postgres?sslmode=disable"
	path_to_migrations = "./migrations"
)

// Config is a description for migrations
type Config struct {
	ConnectionString string
	PathToMigrations string
}

// main Make migrations
func main() {
	cfg := initDB()

	m, err := migrate.New(
		"file://"+cfg.PathToMigrations,
		cfg.ConnectionString,
	)

	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no changes to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}

// Initializing DB
func initDB() *Config {
	var connectionString string
	var pathToMigrations string

	flag.StringVar(&connectionString, "c", "", "connection string")
	flag.StringVar(&pathToMigrations, "m", "", "path to migrations")

	flag.Parse()

	if connectionString == "" {
		panic("connection string is required")
	}

	if pathToMigrations == "" {
		panic("path to migrations is required")
	}

	return &Config{
		ConnectionString: connectionString,
		PathToMigrations: pathToMigrations,
	}
}
