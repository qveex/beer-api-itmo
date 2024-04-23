package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	var dbURL, migrationsPath, migrationsTable string

	flag.StringVar(&dbURL, "db-url", "", "database URL")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migration files")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of the migrations table")
	flag.Parse()
	fmt.Println("Database URL:", dbURL)
	fmt.Println("Migrations Path:", migrationsPath)

	if dbURL == "" {
		fmt.Println("db-url is required")
		os.Exit(1)
	}
	if migrationsPath == "" {
		fmt.Println("migrations-path is required")
		os.Exit(1)
	}

	// Format the migrations path for Windows compatibility
	//formattedPath := "file:///" + strings.ReplaceAll(migrationsPath, "\\", "/")

	m, err := migrate.New(
		"file://"+migrationsPath,
		dbURL,
		//fmt.Sprintf("postgres://%s?x-migrations-table=%s", dbURL, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")
			return
		}
		panic(err)
	}

	fmt.Println("Migrations applied")
}

// Log represents the logger
type Log struct {
	verbose bool
}

// Printf prints out formatted string into a log
func (l *Log) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

// Verbose shows if verbose print enabled
func (l *Log) Verbose() bool {
	return l.verbose
}
