package controller

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/repository"
	"github.com/lilahamstern/ced/server/pkg/logger"
)

type Controller struct {
	repos *repository.Repositories
	log   logger.Logger
}

func New(db *sql.DB) *Controller {
	repos := repository.New(db)
	log := logger.New()
	return &Controller{
		repos,
		log,
	}
}
