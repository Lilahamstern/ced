package service

import "github.com/lilahamstern/ced/server/internal/pkg/repository"

type Services struct {
	Project            ProjectService
	Version            VersionService
	VersionInformation VersionInformationService
}

func New(repos *repository.Repositories) *Services {
	services := &Services{
		Project:            newProject(repos.Project),
		Version:            newVersion(repos.Version, repos.Project),
		VersionInformation: newVersionInformation(repos.VersionInformation),
	}

	return services
}
