package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS : Enable cors for server
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Content-Type", "application/json")

		if c.Request.Method != "OPTIONS" {

			c.Next()
			return
		}

		c.AbortWithStatus(http.StatusOK)
	}
}
