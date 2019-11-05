package database

import (
	"fmt"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"log"
	"os"
)

func (c *Client) Migrate() {
	if os.Getenv("DB_CLEAN_START") == "true" {
		e := c.db.DropTableIfExists(&model.Project{}).Error
		if e != nil {
			log.Panicln(e)
		}
	}

	e := c.db.AutoMigrate(&model.Project{}).Error
	if e != nil {
		fmt.Println("but here")
		log.Panicln(e)
	}

	log.Println("Auto migration succeeded")
}
