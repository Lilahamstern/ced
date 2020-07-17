package v1

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/internal/db/domain"
	"github.com/lilahamstern/ced/server/internal/handler"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model"
	"net/http"
)

func (h *Handler) CreateProject(c *fiber.Ctx) {
	const op errors.Op = "handler.v1.createProject"
	var body model.CreateProject
	// Wants to use custom validation for input data, so no error will be handled
	_ = c.BodyParser(&body)

	if res := body.Validate(); res != nil {
		e := errors.E(op, errors.KindValidation, fmt.Errorf("validation failed"), res)
		c.Next(e)
		return
	}

	err := h.repos.Project.Save(body)
	if err != nil {
		e := errors.E(op, errors.KindInternalServer, err, "Internal server error")
		c.Next(e)
		return
	}

	handler.RespondMessage(c, http.StatusCreated, "Successful creation")
}

func (h *Handler) GetProjects(c *fiber.Ctx) {
	const op errors.Op = "handler.v1.getProjects"

	projects, err := h.repos.Project.GetAll()
	if err != nil {
		e := errors.E(op, errors.KindInternalServer, err, "Internal server error")
		c.Next(e)
		return
	}

	res := map[string]interface{}{"projects": domain.ToModel(projects)}

	handler.RespondData(c, http.StatusOK, res)
}

func (h *Handler) GetProject(c *fiber.Ctx) {
	const op errors.Op = "handler.v1.getProject"
	id := c.Params("id")
	project, err := h.repos.Project.Get(id)
	if err != nil {
		e := errors.E(op, errors.KindInternalServer, err, "Internal Server error")
		c.Next(e)
		return
	}

	if project.ID == 0 {
		handler.RespondMessage(c, http.StatusNotFound, fmt.Sprintf("Project with id %s could not be found.", id))
		return
	}

	versions, err := h.repos.Version.GetVersionsByProjectId(id)
	if err != nil {
		e := errors.E(op, errors.KindInternalServer, err, "Internal Server error")
		c.Next(e)
		return
	}

	project.Versions = versions

	handler.RespondData(c, http.StatusOK, project.ToModel())
}
