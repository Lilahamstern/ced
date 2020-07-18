package handler

import (
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/logger"
	"github.com/lilahamstern/ced/server/pkg/model/response"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func RespondJSON(c *fiber.Ctx, status int, kind string, payload interface{}) {
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

	if err := c.Status(status).JSON(res); err != nil {
		log.Fatalln(err)
	}
}

// ErrorHandler for fiber default error handler
func ErrorHandler(c *fiber.Ctx, err error) {
	code := http.StatusInternalServerError
	msg := "Internal server error"

	e, ok := err.(*errors.Error)
	if !ok {
		RespondJSON(c, code, errors.KindError, msg)
		return
	}

	logger.SystemErr(e)
	RespondJSON(c, e.Status(), e.Kind(), e.Data())
}
