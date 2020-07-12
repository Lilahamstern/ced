package repository

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/repository/project"
)

type Repositories struct {
	Project project.Repository
}

func New(db *sql.DB) *Repositories {
	repos := &Repositories{
		Project: project.NewRepo(db),
	}

	return repos
}
