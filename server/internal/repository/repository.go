package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Project ProjectRepository
	Version VersionRepository
	Component ComponentRepository
}

func New(db *sqlx.DB) *Repositories {
	repos := &Repositories{
		&project{db},
		&version{db},
		&component{db},
	}

	return repos
}
