package main

import (
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
	command            = `go run .\cmd\migrator\ -con-str="postgres://postgres:230704@localhost:5432/postgres?sslmode=disable" -path="./migrations" -op="up"`
)

// Config is a description for migrations
type Config struct {
	ConnectionString string
	PathToMigrations string
	Operation        string
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

	switch cfg.Operation {
	case "down":
		if err = m.Down(); err != nil && err != migrate.ErrNoChange {
			panic(err)
		}
	case "up":
		if err = m.Up(); err != nil {
			panic(err)
		}
	}

	fmt.Println("migrations applied successfully")
}

// Initializing DB
func initDB() *Config {
	var connectionString string
	var pathToMigrations string
	var op string

	flag.StringVar(&connectionString, "con-str", "", "connection string")
	flag.StringVar(&pathToMigrations, "path", "", "path to migrations")
	flag.StringVar(&op, "op", "", "operation")

	flag.Parse()

	if connectionString == "" {
		panic("connection string is required")
	}

	if pathToMigrations == "" {
		panic("path to migrations is required")
	}

	if op == "" {
		panic("operation is required")
	}

	return &Config{
		ConnectionString: connectionString,
		PathToMigrations: pathToMigrations,
		Operation:        op,
	}
}
