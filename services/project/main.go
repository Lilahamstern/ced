package main

import (
	"log"
	"project/config"
	"project/handlers"
	"project/route"
	"project/server"

	_ "github.com/lib/pq"
)

var cfg config.Config

func init() {
	cfg = config.LoadConfig("config.json")
}

func main() {
	handlers.SetupHandler()
	log.Println("Starting service:", cfg.Service.Name)

	app := route.NewRouter(cfg)
	server.StartHTTPServer(cfg.Service.Port, app)
}
