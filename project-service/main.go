package main

import (
	"fmt"
	"github.com/lilahamstern/bec/project-service/service"
	"github.com/lilahamstern/bec/project-service/service/dbclient"
)

var appName = "project-service"

func main() {
	fmt.Printf("Starting %v\n", appName)

	initializeBoltClient()
	service.StartWebService("5050")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
