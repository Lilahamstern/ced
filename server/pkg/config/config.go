package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Load : Loads config
func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Error("Config file could not load: %s", err.Error())
	}
	log.Infof("Config loaded in successfully!")
}

// GenerateDsn : generates database dsn from environment variables
func GenerateDsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%v/%s?%s",
		viper.GetString("database.user"),
		viper.GetString("database.pass"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.flags"),
	)
}
