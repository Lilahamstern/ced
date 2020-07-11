package controller

import (
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/pkg/model"
	"net/http"
)

func (s *Controller) CreateProject(c *fiber.Ctx) {
	var body model.CreateProject
	// Wants to use custom validation for input data, so no error will be handled
	_ = c.BodyParser(&body)

	err := body.Validate()
	if err != nil {
		respondError(c, http.StatusBadRequest, err)
		return
	}

	respondData(c, http.StatusOK, body)
}
