package v1

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/internal/handler"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model"
	"net/http"
)

func (h *Handler) CreateProject(c *fiber.Ctx) {
	const op errors.Op = "controller.createProject"
	var body model.CreateProject
	// Wants to use custom validation for input data, so no error will be handled
	_ = c.BodyParser(&body)

	if res := body.Validate(); res != nil {
		e := errors.E(op, errors.KindValidation, fmt.Errorf("validation failed"), res)
		c.Next(e)
		return
	}

	project, err := h.repos.Project.Save(body)
	if err != nil {
		e := errors.E(op, errors.KindInternalServer, err)
		c.Next(e)
		return
	}

	handler.RespondData(c, http.StatusOK, project.ToModel())
}
