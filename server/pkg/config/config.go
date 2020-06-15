package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"strings"
)

type config struct {
	Port    int
	DbHost  string
	DbPort  int
	DbName  string
	DbUser  string
	DbPass  string
	DbFlags string
}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env config file not found...")
		generateConfigFile()
	}

	log.Println("Loaded config...")
}

func loadDefaultValues() config {
	return config{
		Port:    8080,
		DbHost:  "localhost",
		DbPort:  5432,
		DbName:  "postgres",
		DbUser:  "postgres",
		DbPass:  "postgres",
		DbFlags: "sslmode=disable",
	}
}

func generateConfigFile() {
	log.Println("Generating default .env config...")
	confFile := createConfigFile()
	writeConfigToFile(confFile)

	if err := confFile.Close(); err != nil {
		log.Panicf("Failed to close .env file: %s", err)
	}
	log.Println("Generation complete...")
}

func createConfigFile() *os.File {
	file, err := os.Create(".env")
	if err != nil {
		log.Panicf("Failed to create .env file: %s", err)
	}
	return file
}

func writeConfigToFile(file *os.File) {
	for k, v := range mapValuesFromConfig() {
		data := fmt.Sprintf("%v=%v", k, v)
		if _, err := fmt.Fprintln(file, data); err != nil {
			log.Panicf("Failed writing to .env file: %v", err)
		}
	}
}

func mapValuesFromConfig() map[string]interface{} {
	conf := loadDefaultValues()
	cv := reflect.ValueOf(conf)

	m := make(map[string]interface{})

	for i := 0; i < cv.NumField(); i++ {
		key := strings.ToLower(cv.Type().Field(i).Name)
		value := cv.Field(i).Interface()
		m[key] = value
	}

	return m
}
