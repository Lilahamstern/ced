package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/repository/project"
	"github.com/lilahamstern/ced/server/internal/repository/version"
)

type Repositories struct {
	Project project.Repository
	Version version.Repository
}

func New(db *sqlx.DB) *Repositories {
	repos := &Repositories{
		project.New(db),
		version.New(db),
	}

	return repos
}
