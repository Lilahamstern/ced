package controller

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/pkg/service"
)

type Controller struct {
	services *service.Services
}

func New(db *sql.DB) *Controller {
	services := service.New(db)
	return &Controller{
		services,
	}
}
