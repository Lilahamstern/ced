package repository

import "database/sql"

type Repositories struct {
	Project ProjectRepository
}

func New(db *sql.DB) *Repositories {
	repos := &Repositories{
		Project: newProject(db),
	}

	return repos
}
