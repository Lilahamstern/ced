package main

import (
	"component/config"
	"component/database"
	"component/handlers"
	"component/models"
	"component/route"
	"component/server"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var cfg config.Config

func init() {
	cfg = config.LoadConfig("config.json")
}

func main() {
	handlers.SetupHandler()
	log.Println("Starting service:", cfg.Service.Name)

	db := database.OpenDB(cfg.Database)
	db.Migrate(cfg.Service, &models.Component{})

	app := route.NewRouter(cfg, db)
	server.StartHTTPServer(cfg.Service.Port, app)
}
