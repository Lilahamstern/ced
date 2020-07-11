package service

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/repository"
)

type Services struct {
	Project ProjectService
}

func New(db *sql.DB) *Services {
	repos := repository.New(db)
	services := &Services{
		Project: newProject(repos.Project),
	}

	return services
}
