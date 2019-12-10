package main

import (
	"log"
	"template/config"
	"template/database"
	"template/handlers"
	"template/models"
	"template/route"
	"template/server"

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
	db.Migrate(cfg.Service, &models.AccountData{}, &models.AccountEvent{})

	app := route.NewRouter(cfg, db)
	server.StartHTTPServer(cfg.Service.Port, app)
}
