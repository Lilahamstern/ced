package main

import (
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/ced/server/internal/pkg/server"
	"github.com/lilahamstern/ced/server/pkg/config"
)

func main() {
	conf := config.NewConfig()

	dbSession := database.NewDatabaseConnection(conf)
	dbSession.Migrate()
	conf.AddDbSession(dbSession)

	srv := server.NewHttpServer(conf)

	srv.Start()
	srv.SafeShutDown()
}
