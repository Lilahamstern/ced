package database

import (
	"component/config"
	"log"
)

// Migrate structs into database
func (dc DBClient) Migrate(cfg config.Service, i ...interface{}) {

	if cfg.Clean {
		for _, table := range i {
			if err := dc.db.DropTable(table).Error; err != nil {
				log.Println("Migration error:", err)
			}
		}
		log.Println("Droped tables:", len(i))
	}
	for _, table := range i {
		if err := dc.db.AutoMigrate(table).Error; err != nil {
			log.Println("Migration error:", err)
		}
	}
	log.Println("Migrated tables:", len(i))
}
