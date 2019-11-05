package main

import (
	"fmt"
	"github.com/lilahamstern/bec-microservices/project-service/internal/handlers"
	"github.com/lilahamstern/bec-microservices/project-service/internal/service"
	"github.com/lilahamstern/bec-microservices/project-service/internal/service/dbclient"
)

var appName = "project-service"

func main() {
	fmt.Printf("Starting %v\n", appName)

	initializeMongoClient()
	service.StartWebService("5050")
}

func initializeMongoClient() {
	handlers.DBClient = &dbclient.MongoClient{}
	handlers.DBClient.OpenMongoDb()
}
