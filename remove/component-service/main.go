package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/lilahamstern/bec-microservices/component-service/internal/database"
	"github.com/lilahamstern/bec-microservices/component-service/internal/router"
	"github.com/lilahamstern/bec-microservices/component-service/internal/server"
)

var appName = "component-server"
var dbClient database.IClient

func main() {
	fmt.Printf("Starting %v\n", appName)

	dbClient = &database.Client{}
	dbClient.OpenDb()
	dbClient.Migrate()

	app := router.NewRouter()
	app.Use(dbClient.Inject())
	app.Use(cors.Default())
	router.SetupRouter(app)

	server.StartWebService("5051", app)
}