package main

import (
	"fmt"
	"github.com/lilahamstern/ced/server/internal/repository/database"
	"github.com/lilahamstern/ced/server/internal/server"
	"github.com/lilahamstern/ced/server/pkg/config"
	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	config.Load()
	port := viper.GetInt("port")

	if err := run(port); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(port int) error {
	log := logrus.New()
	log.Out = os.Stdout

	db, dbtidy, err := database.SetupDatabase()
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbtidy()

	s := server.NewServer(db)

	validation.RegisterValidation(db)

	s.Start(port)
	s.SafeShutDown()
	return nil
}
