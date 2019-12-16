package database

import (
	"component/config"
	"fmt"
	"log"

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
