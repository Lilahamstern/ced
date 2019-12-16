package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct
type Config struct {
	Service  `json:"service"`
	Database `json:"database"`
}

// Service struct
type Service struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Port        string `json:"port"`
	Realase     bool   `json:"release"`
	Description string `json:"description"`
	Clean       bool   `json:"clean"`
}

// Database struct
type Database struct {
	Port     string `json:"port"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// LoadConfig loads service config
func LoadConfig(file string) Config {
	var config Config
	cFile, err := os.Open(file)
	if err != nil {
		log.Fatalln(err.Error())
	}

	jsonParser := json.NewDecoder(cFile)
	jsonParser.Decode(&config)
	return config
}
