package space

import (
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/internal/repository"
)

type Handler struct {
	Repos *repository.Repositories
}

func (h *Handler) Routes(space *fiber.Group) {
	space.Post("/projects", h.handleProjectCreate())
	space.Get("/projects", h.handleProjectGetAll())
	space.Get("/projects/:id", h.handleProjectGet())
}
