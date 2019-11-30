package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/lilahamstern/bec-microservices/project-service/internal/database"
	"github.com/lilahamstern/bec-microservices/project-service/internal/router"
	"github.com/lilahamstern/bec-microservices/project-service/internal/server"
)

var appName = "project-server"
var dbClient database.IClient

func main() {
	fmt.Printf("Starting %v\n", appName)

	dbClient = &database.Client{}
	dbClient.OpenDb()
	dbClient.Migrate()

	app := router.NewRouter()
	app.Use(dbClient.Inject())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true;

	app.Use(cors.New(config))
	router.SetupRouter(app)

	server.StartWebService("5050", app)
}
