package repository

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/repository/project"
	"github.com/lilahamstern/ced/server/internal/repository/version"
)

type Repositories struct {
	Project project.Repository
	Version version.Repository
}

func New(db *sql.DB) *Repositories {
	repos := &Repositories{
		project.New(db),
		version.New(db),
	}

	return repos
}
