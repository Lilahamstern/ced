package handlers

import (
	"log"
	"project/route"
)

// SetupHandler for init of handlers
func SetupHandler() {
	log.Println("Init handlers")
	route.NewEndpoint("CreateProject", "POST", "/", create).Add()
	route.NewEndpoint("GetAll", "GET", "/", getAll).Add()
	route.NewEndpoint("Info", "GET", "/info", info).Add()
}
