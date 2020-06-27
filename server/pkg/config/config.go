package config

import (
	"github.com/joho/godotenv"
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"log"
	"os"
	"strconv"
)

// Config struct for general config
type Config struct {
	Port     string
	DB       *database.Session
	Database *DatabaseConfig
	Redis    *RedisConfig
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

// RedisConfig to store conn config
type RedisConfig struct {
	Addr string
	Pass string
	Db   int64
}

// NewConfig
func NewConfig() *Config {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load config: %s", err)
	}

	log.Println("Config loaded...")
	return &Config{
		Port:     getEnv("port"),
		Database: newDatabaseConfig(),
		Redis:    newRedisConfig(),
	}
}

func (c *Config) AddDbSession(db *database.Session) {
	c.DB = db
}

func newDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Port:  getEnvInt("dbport"),
		Host:  getEnv("dbhost"),
		Name:  getEnv("dbname"),
		User:  getEnv("dbuser"),
		Pass:  getEnv("dbpass"),
		Flags: getEnv("dbflag"),
	}
}

func newRedisConfig() *RedisConfig {
	return &RedisConfig{
		Addr: getEnv("redisaddr"),
		Pass: getEnv("redispass"),
		Db:   getEnvInt("redisdb"),
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
