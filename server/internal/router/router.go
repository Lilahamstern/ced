package router

import (
	"github.com/gofiber/fiber"
	"github.com/jmoiron/sqlx"
	v1 "github.com/lilahamstern/ced/server/internal/handler/v1"
	"github.com/lilahamstern/ced/server/internal/repository"
)

type Router struct {
	repos *repository.Repositories
}

func New(db *sqlx.DB) *Router {
	repos := repository.New(db)
	return &Router{
		repos,
	}
}

func (h *Router) RegisterRoutes(app *fiber.App) {
	v1Handler := v1.New(h.repos)

	api := app.Group("/api")

	v1Handler.RegisterRoutes(api)
}
