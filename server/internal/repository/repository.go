package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Project ProjectRepository
	Version VersionRepository
}

func New(db *sqlx.DB) *Repositories {
	repos := &Repositories{
		&project{db},
		&version{db},
	}

	return repos
}
