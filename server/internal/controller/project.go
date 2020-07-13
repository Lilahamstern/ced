package controller

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/logger"
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

	exists, err := s.repos.Project.Exists(body.ID)
	if err != nil {
		e := errors.E(op, errors.KindAlreadyExists, err)
		c.Next(e)
	}

	if exists {
		// TODO: Add exist validation to validator
		msg := fmt.Sprintf("project already exists with id: %v", body.ID)
		respondError(c, http.StatusConflict, msg)
		return
	}

	project, err := s.repos.Project.Save(body)
	if err != nil {
		logger.SystemErr(err)
		respondError(c, errors.Kind(err), "Internal Server Error")
		return
	}

	respondData(c, http.StatusOK, project)
}
