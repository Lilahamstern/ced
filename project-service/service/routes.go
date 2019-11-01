package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{
		Name:        "Getproject",
		Method:      "GET",
		Pattern:     "/projects/{projectId}",
		HandlerFunc: GetProject,
	},
}
