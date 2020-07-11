package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Config struct for general config
type Config struct {
	Port     string
	Database *DatabaseConfig
}

// DatabaseConfig to store conn config
type DatabaseConfig struct {
	Port  int64
	Host  string
	Name  string
	User  string
	Pass  string
	Flags string
}

// New
func New() *Config {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load config: %s", err)
	}

	log.Println("Config loaded...")
	return &Config{
		Port:     getEnv("port"),
		Database: newDatabaseConfig(),
	}
}

// generateDbUrl generates database url from environment variables
// Returns database url as string
func (c Config) GenerateDbUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%v/%s?%s",
		c.Database.User,
		c.Database.Pass,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.Flags,
	)
}

func newDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Port:  getEnvInt("dbport"),
		Host:  getEnv("dbhost"),
		Name:  getEnv("dbname"),
		User:  getEnv("dbuser"),
		Pass:  getEnv("dbpass"),
		Flags: getEnv("dbflags"),
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func getEnvInt(key string) int64 {
	env, err := strconv.ParseInt(getEnv(key), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse %s to int", key)
	}

	return env
}
