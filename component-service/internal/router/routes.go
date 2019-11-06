package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/component-service/internal/handlers"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{
		Method:      "GET",
		Pattern:     "/components/:projectId",
		HandlerFunc: handlers.GetComponents,
	},
	{
		Method:      "POST",
		Pattern:     "/components/:projectId",
		HandlerFunc: handlers.CreateComponents,
	},
	{
		Method:      "GET",
		Pattern:     "/health",
		HandlerFunc: handlers.HealthCheck,
	},
}
