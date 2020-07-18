package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lilahamstern/ced/server/pkg/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

type database struct {
	DB *sqlx.DB
}

func SetupDatabase() (*sqlx.DB, func(), error) {
	dsn := config.GenerateDsn()

	db, err := connect(dsn)
	if err != nil {
		return nil, nil, err
	}
	log.Println("Connected to database...")

	if err = db.Migrate(); err != nil {
		return nil, nil, err
	}
	log.Println("Migrated database...")

	tidy := func() {
		if err := db.DB.Close(); err != nil {
			log.Errorf("Could not close database connection: %s\n", err)
		}
	}

	return db.DB, tidy, nil
}

func connect(dsn string) (*database, error) {
	var db *sqlx.DB
	var err error

	for i := 1; i < 8; i++ {
		log.Printf("Trying to connect to the database (attempt %d)\n", i)

		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}

		log.Printf("Error: %v\n", err)

		time.Sleep(time.Duration(i*i) * time.Second)
	}

	if db == nil {
		return nil, errors.Wrap(err, "Failed to create connection to database")
	}

	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)

	return &database{db}, nil
}

// Migrate current database migrations to db. If they aren't applied
func (db *database) Migrate() error {
	driver, err := postgres.WithInstance(db.DB.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/repository/database/migrations/postgres",
		"postgres",
		driver,
	)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "Migration failed")
	}
	return nil
}
