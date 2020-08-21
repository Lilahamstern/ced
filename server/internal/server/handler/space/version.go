package space

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/internal/server/handler"
	. "github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model/request"
)

func (h *Handler) handleVersionCreate() gin.HandlerFunc {
	const op Op = "handler.space.handleVersionCreate"

	type response struct {
		Message string `json:"message"`
	}

	return func(c *gin.Context) {
		projectID, _ := c.Params.Get("id")
		var body request.CreateVersion

		err := c.BindJSON(&body)
		if err != nil {
			e := E(op, err, KindError)
			c.Error(e)
			c.Next()
			return
		}

		if res := body.Validate(); res != nil {
			e := E(op, http.StatusBadRequest, KindFail, errors.New("validation failed"), res)
			c.Error(e)
			c.Next()
			return
		}

		err = h.Repos.Version.Save(projectID, body)
		if err != nil {
			e := E(op, err, KindError)
			c.Error(e)
			c.Next()
			return
		}

		res := response{
			"Version created",
		}
		handler.RespondJSON(c, http.StatusCreated, KindSuccess, res)
	}
}
