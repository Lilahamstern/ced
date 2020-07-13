package controller

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/repository"
)

type Controller struct {
	repos *repository.Repositories
}

func New(db *sql.DB) *Controller {
	repos := repository.New(db)
	return &Controller{
		repos,
	}
}
