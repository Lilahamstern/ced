package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/project-service/internal/handlers"
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
		Pattern:     "/projects/:projectId",
		HandlerFunc: handlers.GetProject,
	},
	{
		Method:      "GET",
		Pattern:     "/projects",
		HandlerFunc: handlers.GetAllProjects,
	},
	{
		Method:      "POST",
		Pattern:     "/projects",
		HandlerFunc: handlers.CreateProject,
	},
	{
		Method:      "DELETE",
		Pattern:     "/projects/:projectId",
		HandlerFunc: handlers.DeleteProject,
	},
	{
		Method:      "GET",
		Pattern:     "/health",
		HandlerFunc: handlers.HealthCheck,
	},
}
