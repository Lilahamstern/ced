package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

func GenerateConfig(config Config) {
	conf := createConfigFile()
	writeConfigFile(conf)
}

func writeConfigFile(wr io.Writer) {
	size, err := fmt.Fprintln(wr, "heyy")
	if err != nil {
		log.Fatalf("Failed writing to config: %s", err)
	}

	fmt.Println(size)
}

func createConfigFile() *os.File {
	file, err := os.Create(".env")
	if err != nil {
		log.Fatalf("Could not create .env file: %v", err.Error())
	}

	return file
}

func mapValuesFromConfig(config Config) []string {

}

//
//// generateConfigFil generates config file depending on config struct
//func generateConfigFile() {
//	log.Println("Generating default .env config...")
//	confFile := createConfigFile()
//	writeConfigToFile(confFile)
//
//	if err := confFile.Close(); err != nil {
//		log.Panicf("Failed to close .env file: %s", err)
//	}
//	log.Println("Generation complete...")
//}
//
//// createConfigFile will creating config file and return *os.File
//func createConfigFile() *os.File {
//	file, err := os.Create(".env")
//	if err != nil {
//		log.Panicf("Failed to create .env file: %s", err)
//	}
//	return file
//}
//
//// writeConfigToFile will write config and its values to .env file
//func writeConfigToFile(file *os.File) {
//	for k, v := range mapValuesFromConfig() {
//		data := fmt.Sprintf("%v=%v", k, v)
//		if _, err := fmt.Fprintln(file, data); err != nil {
//			log.Panicf("Failed writing to .env file: %v", err)
//		}
//	}
//}
//
//// mapValuesFromConfig takes the loaded config and put it into a map[string]interface{}
//// It returns the loaded config in map with key:value
//func mapValuesFromConfig() map[string]interface{} {
//	conf := loadDefaultValues()
//	cv := reflect.ValueOf(conf)
//
//	m := make(map[string]interface{})
//
//	for i := 0; i < cv.NumField(); i++ {
//		key := strings.ToLower(cv.Type().Field(i).Name)
//		value := cv.Field(i).Interface()
//		m[key] = value
//	}
//
//	return m
//}
