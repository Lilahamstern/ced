package main

import (
	"fmt"
	"github.com/lilahamstern/ced/server/pkg/logger"
	"github.com/lilahamstern/ced/server/pkg/logger/sentry"
	"os"

	"github.com/lilahamstern/ced/server/internal/repository/database"
	"github.com/lilahamstern/ced/server/internal/server"
	"github.com/lilahamstern/ced/server/pkg/config"
	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/pkg/errors"
)

func main() {
	config.Load()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger.Register()

	sentrytidy, err := sentry.Register()
	if err != nil {
		return errors.Wrap(err, "Sentry register")
	}
	defer sentrytidy()

	db, dbtidy, err := database.SetupDatabase()
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbtidy()

	s := server.NewServer(db)

	validation.RegisterValidation(db)

	s.Start()
	return nil
}
