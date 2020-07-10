package repository

import "database/sql"

type Repositories struct {
	Project            ProjectRepository
	Version            VersionRepository
	VersionInformation VersionInformationRepository
}

func New(db *sql.DB) *Repositories {
	repos := &Repositories{
		Project:            newProject(db),
		Version:            newVersion(db),
		VersionInformation: newVersionInformation(db),
	}

	return repos
}
