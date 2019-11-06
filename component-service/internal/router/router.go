package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func SetupRouter(engine *gin.Engine) {
	for _, route := range routes {
		engine.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}
