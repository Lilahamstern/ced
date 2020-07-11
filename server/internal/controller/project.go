package controller

import (
	"github.com/gofiber/fiber"
	"net/http"
)

func (s *Controller) CreateProject(c *fiber.Ctx) {
	c.Status(http.StatusOK).JSON("Created project")
}
