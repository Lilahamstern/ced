package route

import (
	"strings"
	"template/config"
	"template/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Route struct
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// Routes struct
var routes = make(map[string]Route)

// NewRouter creates new router instance and initilizing the router context
func NewRouter(cfg config.Config, db database.DBClient) *gin.Engine {
	if cfg.Service.Realase {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(inject(cfg.Service, db))

	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	router.Use(cors.New(conf))

	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	return router
}

func inject(cfg config.Service, db database.DBClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("cfg", cfg)
		ctx.Set("db", db)
		ctx.Next()
	}
}

// Add route to endpoint
func (r Route) Add() Route {
	routes[strings.ToLower(r.Path)] = r
	return r
}

// NewEndpoint getting added to the engine
func NewEndpoint(method string, path string, handler gin.HandlerFunc) Route {
	return Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
