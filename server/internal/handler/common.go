package handler

import (
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/logger"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func respondJSON(c *fiber.Ctx, status int, payload interface{}) {
	if err := c.Status(status).JSON(payload); err != nil {
		log.Fatalln(err)
	}
}

func RespondData(c *fiber.Ctx, status int, payload interface{}) {
	respondJSON(c, status, map[string]interface{}{"data": payload})
}

func RespondError(c *fiber.Ctx, status int, payload interface{}) {
	respondJSON(c, status, map[string]interface{}{"errors": payload})
}

// ErrorHandler for fiber default error handler
func ErrorHandler(c *fiber.Ctx, err error) {
	code := http.StatusInternalServerError
	msg := "Internal server error"

	e, ok := err.(*errors.Error)
	if !ok {
		RespondError(c, code, msg)
	}

	logger.SystemErr(e)
	RespondError(c, errors.Kind(e), e.Data())
}
