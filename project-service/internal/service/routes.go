package service

import (
	"github.com/lilahamstern/bec-microservices/project-service/internal/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{
		Name:        "GetProject",
		Method:      "GET",
		Pattern:     "/projects/{projectId}",
		HandlerFunc: handlers.GetProject,
	},
	{
		Name:        "CreateProject",
		Method:      "POST",
		Pattern:     "/projects/",
		HandlerFunc: handlers.CreateProject,
	},
	{
		Name:        "HealthCheck",
		Method:      "GET",
		Pattern:     "/health",
		HandlerFunc: handlers.HealthCheck,
	},
}
