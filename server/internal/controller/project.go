package controller

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model"
	"net/http"
)

func (s *Controller) CreateProject(c *fiber.Ctx) {
	const op errors.Op = "controller.createProject"
	var body model.CreateProject
	// Wants to use custom validation for input data, so no error will be handled
	_ = c.BodyParser(&body)

	if res := body.Validate(); res != nil {
		e := errors.E(op, errors.KindValidation, fmt.Errorf("validation failed"), res)
		c.Next(e)
		return
	}

	project, err := s.repos.Project.Save(body)
	if err != nil {
		e := errors.E(op, errors.KindInternalServer, err)
		c.Next(e)
		return
	}

	respondData(c, http.StatusOK, project)
}
