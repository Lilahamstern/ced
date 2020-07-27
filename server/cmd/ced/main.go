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

// @title CED Rest API
// @version 0.1
// @description Offical CED Rest API docs.
// @termsOfService http://swagger.io/terms/

// @contact.name Author
// @contact.url http://www.hamsterapps.net
// @contact.email leo.ronnebro@hamsterapps.net

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
