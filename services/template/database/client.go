package database

import (
	"fmt"
	"log"
	"template/config"
	"template/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBClient struct
type DBClient struct {
	db *gorm.DB
}

// OpenDB connection
func OpenDB(cfg config.Database) DBClient {
	var err error
	var dc DBClient

	// pgInfo := fmt.Sprintf("host=%s port=%v user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	dc.db, err = gorm.Open("postgres", fmt.Sprintf("postgresql://%v:%v@%v/%v?sslmode=disable", cfg.User, cfg.Password, cfg.Host+":"+cfg.Port, cfg.Name))
	if err != nil {
		log.Fatal("Could not connect to db:", err)
	}

	dc.db.LogMode(false)

	dc.db.AutoMigrate(&models.AccountData{}, &models.AccountEvent{})

	return dc
}

// Check database status
func (dc *DBClient) Check() bool {
	return dc.db != nil
}
