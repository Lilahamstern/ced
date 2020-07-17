package v1

import (
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/internal/repository"
)

type Handler struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *Handler {
	return &Handler{
		repos,
	}
}

func (h *Handler) RegisterRoutes(api *fiber.Group) {
	v1 := api.Group("/v1")
	{
		v1.Post("/projects", h.CreateProject)
		v1.Get("/projects", h.GetProjects)
		v1.Get("/projects/:id", h.GetProject)
	}
}
