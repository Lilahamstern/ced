package middleware

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/internal/server/handler"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/logger"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			e, ok := err.Err.(*errors.Error)
			if !ok {
				logger.Error(e)
			}

			status := e.Code()
			if status == 0 {
				status = http.StatusInternalServerError
			}

			logger.SystemErr(e)

			switch e.Status() {
			case errors.KindError:
				eId := sentry.CaptureException(e)
				handler.RespondJSON(c, status, e.Status(), fmt.Sprintf("Internal server error, ID: %s", eId))
			default:
				handler.RespondJSON(c, status, e.Status(), e.Data())
			}
		}
	}
}
