package main

import (
	database "github.com/lilahamstern/ced/server/internal/db/postgres"
	"github.com/lilahamstern/ced/server/internal/repository"
	"github.com/lilahamstern/ced/server/internal/resolver"
	"github.com/lilahamstern/ced/server/internal/server"
	"github.com/lilahamstern/ced/server/pkg/config"
	"github.com/lilahamstern/ced/server/pkg/service"
)

func main() {
	conf := config.NewConfig()

	dbSession := database.NewConnection(conf)
	dbSession.Migrate()

	srv := server.NewHttpServer(conf)

	registerResolvers(dbSession)

	srv.Start()
	srv.SafeShutDown()
}

func registerResolvers(db *database.Session) {
	repos := repository.New(db.DB)

	services := service.New(repos)

	resolver.NewResolver(services)
}
