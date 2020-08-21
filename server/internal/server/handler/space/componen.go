package space

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/internal/server/handler"
	. "github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model/request"
)

func (h *Handler) handleComponentGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		component := request.CreateComponent{}
		handler.RespondJSON(c, 200, KindSuccess, component)
	}
}

func (h *Handler) handleComponentCreate() gin.HandlerFunc {
	const op Op = "handler.space.handleComponentCreate"

	type response struct {
		Message string `json:"message"`
	}

	return func(c *gin.Context) {
		versionID, _ := c.Params.Get("id")
		var body []request.CreateComponent

		err := c.BindJSON(&body)
		if err != nil {
			e := E(op, err, KindError)
			c.Error(e)
			c.Next()
			return
		}

		clean := []*request.CreateComponent{}
		for _, component := range body {
			if res := component.Validate(); res != nil {
				e := E(op, http.StatusBadRequest, KindFail, errors.New("validation failed"), res)
				c.Error(e)
				c.Next()
				return
			}

			clean = append(clean, &component)
		}

		err = h.Repos.Component.Save(clean)
		if err != nil {
			e := E(op, err, KindError)
			c.Error(e)
			c.Next()
			return
		}

		handler.RespondJSON(c, 400, KindSuccess, versionID)
		// if res := body.Validate(); res != nil {
		// 	e := E(op, http.StatusBadRequest, KindFail, errors.New("validation failed"), res)
		// 	c.Error(e)
		// 	c.Next()
		// 	return
		// }

		// err = h.Repos.Version.Save(projectID, body)
		// if err != nil {
		// 	e := E(op, err, KindError)
		// 	c.Error(e)
		// 	c.Next()
		// 	return
		// }

		// res := response{
		// 	"Version created",
		// }
		// handler.RespondJSON(c, http.StatusCreated, KindSuccess, res)
	}
}
