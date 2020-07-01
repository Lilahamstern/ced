package resolver

import (
	"github.com/lilahamstern/ced/server/pkg/service"
)

//go:generate go run github.com/99designs/gqlgen

var resolver *Resolver

type Resolver struct {
	ProjectService            service.ProjectService
	VersionService            service.VersionService
	VersionInformationService service.VersionInformationService
}

func NewResolver(projectService service.ProjectService, versionService service.VersionService, versionInformationService service.VersionInformationService) {
	resolver = &Resolver{
		projectService,
		versionService,
		versionInformationService,
	}
}
