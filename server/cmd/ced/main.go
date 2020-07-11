package main

import (
	"github.com/lilahamstern/ced/server/internal/controller"
	database "github.com/lilahamstern/ced/server/internal/db/postgres"
	"github.com/lilahamstern/ced/server/internal/server"
	"github.com/lilahamstern/ced/server/pkg/config"
)

func main() {
	conf := config.New()

	dbSession := database.NewConnection(conf)
	dbSession.Migrate()

	controllers := controller.New(dbSession.DB)
	srv := server.New(conf.Port, controllers)

	srv.Start()
	srv.SafeShutDown()
}
