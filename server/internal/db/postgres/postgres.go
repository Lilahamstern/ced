package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/lilahamstern/ced/server/pkg/config"
	"github.com/lilahamstern/ced/server/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

type Session struct {
	DB *sql.DB
}

const connectionNotEstablishedError = "Database connection is not establishd: %v"

// NewConnection establish database connection depending on config
// If connection fails reconnection will be tried maximum of 5 times
func NewConnection(config *config.Config) *Session {

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

	if err := session.PingCheck(); err != nil {
		log.Fatalf(connectionNotEstablishedError, err)
	}

	log.Println("Connected to database...")

	return session
}

// PingCheck ping database if that does not succeed error will be thrown.
func (s *Session) PingCheck() error {
	if err := s.DB.Ping(); err != nil {
		return err
	}
	return nil
}

// Migrate current database migrations to db. If they aren't applied
func (s *Session) Migrate() {
	if err := s.PingCheck(); err != nil {
		log.Fatalf(connectionNotEstablishedError, err)
	}

	driver, err := postgres.WithInstance(s.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations/postgres",
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

func (s *Session) UniqueRecord(table string, field string, data interface{}) (bool, error) {
	const op errors.Op = "postgres.uniqueRecord"
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s=$1)", table, field)

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return false, errors.E(op, err, log.WarnLevel)
	}
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRow(data).Scan(&exists)
	if err != nil {
		return false, errors.E(op, err, log.WarnLevel)
	}

	return exists, nil
}
