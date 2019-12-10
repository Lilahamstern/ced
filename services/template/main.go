package main

import (
	"log"
	"template/config"
	"template/database"
	"template/handlers"
	"template/route"
	"template/server"
)

var cfg config.Config

func init() {
	cfg = config.LoadConfig("config.json")
}

func main() {
	handlers.SetupHandler()
	log.Println("Starting service:", cfg.Service.Name)

	app := route.NewRouter(cfg, database.OpenDB(cfg.Database))

	server.StartHTTPServer(cfg.Service.Port, app)
}
