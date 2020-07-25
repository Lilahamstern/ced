package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model/response"
)

func RespondJSON(c *gin.Context, code int, status errors.Status, payload interface{}) {
	res := response.Response{
		Status: status.String(),
	}

	switch payload.(type) {
	case string:
		res.Message = payload.(string)
	case error:
		res.Message = payload.(string)
	default:
		res.Data = payload
	}

	c.JSON(code, res)
}
