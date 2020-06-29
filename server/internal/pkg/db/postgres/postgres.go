package database

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/lilahamstern/ced/server/pkg/config"
	"log"
	"time"
)

type Session struct {
	DB *sql.DB
}

// InitDB will initialize database connection
// If connection fails reconnection will be tried
func NewDatabaseConnection(config *config.Config) *Session {

	conn := config.GenerateDbUrl()
	var db *sql.DB
	var err error
	for i := 1; i < 5; i++ {
		log.Printf("Trying to connect to the database (attempt %d)...\n", i)

		db, err = sql.Open("postgres", conn)
		if err == nil {
			break
		}
		log.Printf("Error: %v\n", err)

		time.Sleep(time.Duration(i*i) * time.Second)
	}

	if db == nil {
		log.Fatalln("could not connect to database...")
	}

	session := &Session{
		DB: db,
	}

	session.PingCheck()
	log.Println("Connected to database...")

	return session
}

// PingCheckDB ping database if that does not succeed error will be thrown.
func (s *Session) PingCheck() {
	if err := s.DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

// Migrate current database migrations to db. If they arent in the database
func (s *Session) Migrate() {
	s.PingCheck()
	driver, err := postgres.WithInstance(s.DB, &postgres.Config{})
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
