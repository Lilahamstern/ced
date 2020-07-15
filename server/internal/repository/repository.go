package repository

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/repository/information"
	"github.com/lilahamstern/ced/server/internal/repository/project"
)

type Repositories struct {
	Project     project.Repository
	Information information.Repository
}

func New(db *sql.DB) *Repositories {
	repos := &Repositories{
		Project:     project.New(db),
		Information: information.New(db),
	}

	return repos
}
