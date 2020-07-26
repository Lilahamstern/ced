package main

import (
	"fmt"
	"os"

	"github.com/lilahamstern/ced/server/pkg/logger"
	"github.com/lilahamstern/ced/server/pkg/logger/sentry"

	"github.com/lilahamstern/ced/server/internal/repository/database"
	"github.com/lilahamstern/ced/server/internal/server"
	"github.com/lilahamstern/ced/server/pkg/config"
	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/pkg/errors"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /space
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	err := config.Load()
	if err != nil {
		return errors.Wrap(err, "Config load")
	}

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
