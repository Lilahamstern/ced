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

	router.Use(inject("cfg", cfg))
	router.Use(inject("db", db))

	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	router.Use(cors.New(conf))

	return router
}

func inject(name string, c ...interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(name, c)
		ctx.Next()
	}
}

// Add route to endpoint
func (r Route) Add() Route {
	routes[strings.ToLower(r.Path)] = r
	return r
}

// AddHandler to engine
func AddHandler(method string, path string, handler gin.HandlerFunc) {

}
