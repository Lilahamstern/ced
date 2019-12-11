package database

import (
	"fmt"
	"log"
	"project/config"

	"github.com/jinzhu/gorm"
)

//DBClient struct
type DBClient struct {
	db *gorm.DB
}

// OpenDB connection
func OpenDB(cfg config.Database) DBClient {
	var err error
	var dc DBClient

	dc.db, err = gorm.Open("postgres", fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", cfg.User, cfg.Password, cfg.Host+":"+cfg.Port, cfg.Name))
	if err != nil {
		log.Fatal("Coud not connect to db:", err)
	}

	dc.db.LogMode(false)

	return dc
}

// Check database status
func (dc *DBClient) Check() bool {
	return dc.db != nil
}

// addExtensions adds provided exts
func (dc *DBClient) addExtensions(exts []string) {
	for _, v := range exts {
		query := fmt.Sprintf(`CREATE EXTENSION IF NOT EXISTS "%s" WITH SCHEMA public;`, v)
		if err := dc.db.Exec(query).Error; err != nil {
			log.Println(err)
		}
		log.Printf("[DB] Added extention %s", v)
	}
}
