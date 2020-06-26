package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Config struct for general config
type Config struct {
	Port     int64
	database *DatabaseConfig
	redis    *RedisConfig
}

// DatabaseConfig to store conn config
type DatabaseConfig struct {
	port  int64
	host  string
	name  string
	user  string
	pass  string
	flags string
}

// RedisConfig to store conn config
type RedisConfig struct {
	addr string
	pass string
	db   int64
}

// NewConfig
func NewConfig() *Config {

	if err := godotenv.Load(); err != nil {
		GenerateConfig(Config{})
	}

	log.Println("Config loaded...")
	return &Config{
		Port:     getEnvInt("port"),
		database: newDatabaseConfig(),
		redis:    newRedisConfig(),
	}
}

func newDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		port:  getEnvInt("dbport"),
		host:  getEnv("dbhost"),
		name:  getEnv("dbname"),
		user:  getEnv("dbuser"),
		pass:  getEnv("dbpass"),
		flags: getEnv("dbflag"),
	}
}

func newRedisConfig() *RedisConfig {
	return &RedisConfig{
		addr: getEnv("redisaddr"),
		pass: getEnv("redispass"),
		db:   getEnvInt("redisdb"),
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
