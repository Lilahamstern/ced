package main

import (
	"github.com/lilahamstern/ced/server/internal/controller"
	database "github.com/lilahamstern/ced/server/internal/db/postgres"
	"github.com/lilahamstern/ced/server/internal/server"
	"github.com/lilahamstern/ced/server/pkg/config"
	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	log := logrus.New()
	log.Out = os.Stdout

	conf := config.New()

	dbSession := database.NewConnection(conf)
	dbSession.Migrate()

	validation.Register(dbSession)
	controllers := controller.New(dbSession.DB)
	srv := server.New(conf.Port, controllers)

	srv.Start()
	srv.SafeShutDown()
}
