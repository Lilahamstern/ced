package main

import (
	database "github.com/lilahamstern/ced/server/internal/db/postgres"
	"github.com/lilahamstern/ced/server/internal/router"
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
	routes := router.New(dbSession.DB)
	srv := server.New(conf.Port, routes, log.Out)

	srv.Start()
	srv.SafeShutDown()
}
