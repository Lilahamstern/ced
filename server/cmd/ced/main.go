package main

import (
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/ced/server/internal/pkg/repository"
	"github.com/lilahamstern/ced/server/internal/pkg/server"
	"github.com/lilahamstern/ced/server/internal/resolver"
	"github.com/lilahamstern/ced/server/pkg/config"
	"github.com/lilahamstern/ced/server/pkg/service"
)

func main() {
	conf := config.NewConfig()

	dbSession := database.NewDatabaseConnection(conf)
	dbSession.Migrate()

	srv := server.NewHttpServer(conf)

	registerResolvers(dbSession)

	srv.Start()
	srv.SafeShutDown()
}

func registerResolvers(db *database.Session) {
	projectRepository := repository.NewProjectRepository(db.DB)
	versionRepository := repository.NewVersionRepository(db.DB)
	versionInfoRepository := repository.NewVersionInformationRepository(db.DB)

	projectService := service.NewProjectService(projectRepository)
	versionService := service.NewVersionService(versionRepository, projectRepository)
	versionInformationService := service.NewVersionInformationService(versionInfoRepository)

	resolver.NewResolver(projectService, versionService, versionInformationService)
}
