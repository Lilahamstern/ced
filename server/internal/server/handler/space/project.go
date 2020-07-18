package space

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/internal/repository/domain"
	"github.com/lilahamstern/ced/server/internal/server/handler"
	. "github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model/request"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) handleProjectCreate() fiber.Handler {
	const op Op = "handler.space.handleProjectCreate"

	type response struct {
		Message string `json:"message"`
	}

	var body request.CreateProject
	return func(c *fiber.Ctx) {
		_ = c.BodyParser(&body)

		if res := body.Validate(); res != nil {
			e := E(op, http.StatusBadRequest, KindFail, errors.New("validation failed"), res)
			c.Next(e)
			return
		}

		err := h.Repos.Project.Save(body)
		if err != nil {
			e := E(op, http.StatusInternalServerError, err, "Internal server error")
			c.Next(e)
			return
		}

		res := response{
			"Project created",
		}
		handler.RespondJSON(c, http.StatusCreated, KindSuccess, res)
	}
}

func (h *Handler) handleProjectGetAll() fiber.Handler {
	const op Op = "handler.space.handleProjectGetAll"

	type Response struct {
		Projects interface{} `json:"projects"`
	}

	return func(c *fiber.Ctx) {
		projects, err := h.Repos.Project.GetAll()
		if err != nil {
			e := E(op, err, "Internal server error")
			c.Next(e)
			return
		}

		res := Response{
			domain.ToModel(projects),
		}

		handler.RespondJSON(c, http.StatusOK, KindSuccess, res)
	}
}

func (h *Handler) handleProjectGet() fiber.Handler {
	const op Op = "handler.space.handleProjectGet"

	type Response struct {
		Project interface{} `json:"project"`
	}
	return func(c *fiber.Ctx) {
		id := c.Params("id")

		project, err := h.Repos.Project.Get(id)
		if err != nil {
			e := E(op, err, "Internal Server error")
			c.Next(e)
			return
		}

		if project == nil {
			message := fmt.Sprintf("Project with id %s could not be found.", id)
			handler.RespondJSON(c, http.StatusNotFound, KindFail, message)
			return
		}

		versions, err := h.Repos.Version.GetVersionsByProjectId(id)
		if err != nil {
			e := E(op, err, "Internal Server error")
			c.Next(e)
		}

		project.Versions = versions

		res := &Response{
			project.ToModel(),
		}

		handler.RespondJSON(c, http.StatusOK, KindSuccess, res)
	}
}
