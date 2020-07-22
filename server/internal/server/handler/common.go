package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/pkg/model/response"
)

func RespondJSON(c *gin.Context, status int, kind string, payload interface{}) {
	res := response.Response{
		Kind: kind,
	}

	switch payload.(type) {
	case string:
		res.Message = payload.(string)
	case error:
		res.Message = payload.(string)
	default:
		res.Data = payload
	}

	c.JSON(status, res)
}
