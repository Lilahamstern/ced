package database

import (
	"github.com/lilahamstern/bec-microservices/component-service/internal/model"
	"log"
	"os"
)

func (c *Client) Migrate() {
	if os.Getenv("DB_CLEAN_START") == "true" {
		e := c.db.DropTableIfExists(&model.Component{}).Error
		if e != nil {
			log.Panicln(e)
		}
	}

	e := c.db.AutoMigrate(&model.Component{}).Error
	if e != nil {
		log.Panicln(e)
	}

	log.Println("Auto migration succeeded")
}
