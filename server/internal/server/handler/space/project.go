package space

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/internal/repository/domain"
	"github.com/lilahamstern/ced/server/internal/server/handler"
	. "github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model/request"
	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) handleProjectCreate() gin.HandlerFunc {
	const op Op = "handler.space.handleProjectCreate"

	type response struct {
		Message string `json:"message"`
	}

	var body request.CreateProject
	return func(c *gin.Context) {
		_ = c.BindJSON(&body)

		if res := body.Validate(); res != nil {
			e := E(op, http.StatusBadRequest, KindFail, errors.New("validation failed"), res)
			c.Error(e)
			c.Next()
			return
		}

		err := h.Repos.Project.Save(body)
		if err != nil {
			e := E(op, err, KindError)
			c.Error(e)
			c.Next()
			return
		}

		res := response{
			"Project created",
		}
		handler.RespondJSON(c, http.StatusCreated, KindSuccess, res)
	}
}

func (h *Handler) handleProjectGetAll() gin.HandlerFunc {
	const op Op = "handler.space.handleProjectGetAll"

	type Response struct {
		Projects interface{} `json:"projects"`
	}

	return func(c *gin.Context) {
		projects, err := h.Repos.Project.GetAll()
		if err != nil {
			e := E(op, err, KindError)
			c.Error(e)
			c.Next()
			return
		}

		res := Response{
			domain.ToModel(projects),
		}

		handler.RespondJSON(c, http.StatusOK, KindSuccess, res)
	}
}

func (h *Handler) handleProjectGet() gin.HandlerFunc {
	const op Op = "handler.space.handleProjectGet"

	type Response struct {
		Project interface{} `json:"project"`
	}
	return func(c *gin.Context) {
		id := c.Param("id")

		project, err := h.Repos.Project.Get(id)
		if err != nil {
			e := E(op, err, KindError)
			c.Error(e)
			c.Next()
			return
		}

		if project.ID == 0 {
			message := fmt.Sprintf("Project with id %s could not be found.", id)
			e := E(op, http.StatusNotFound, KindFail, message)
			c.Error(e)
			c.Next()
			return
		}

		versions, err := h.Repos.Version.GetVersionsByProjectId(id)
		if err != nil {
			e := E(op, err, KindError)
			c.Error(e)
			c.Next()
		}

		project.Versions = versions

		res := &Response{
			project.ToModel(),
		}

		handler.RespondJSON(c, http.StatusOK, KindSuccess, res)
	}
}
