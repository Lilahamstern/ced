package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

var DB *sql.DB

// InitDB will initialize database connection
// If connection fails reconnection will be tried
func InitDB() {

	var db *sql.DB
	var err error
	for i := 1; i < 5; i++ {
		log.Printf("Trying to connect to the database (attempt %d)...\n", i)
		db, err = sql.Open("postgres", generateDbUrl())
		if err == nil {
			break
		}
		log.Printf("Error: %v\n", err)

		time.Sleep(time.Duration(i*i) * time.Second)
	}

	if db == nil {
		log.Fatalln("could not connect to database...")
	}

	DB = db

	PingCheckDB()
	log.Println("Connected to database...")
}

// PingCheckDB ping database if that does not succeed error will be thrown.
func PingCheckDB() {
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

// Migrate current database migrations to db. If they arent in the database
func Migrate() {
	PingCheckDB()
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/postgres",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to migrate db: %s", err)
	}
	log.Println("Migrated database...")
}

// generateDbUrl generates database url from environment variables
// Returns database url as string
func generateDbUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		os.Getenv("dbuser"),
		os.Getenv("dbpass"),
		os.Getenv("dbhost"),
		os.Getenv("dbport"),
		os.Getenv("dbname"),
		os.Getenv("dbflags"))
}
