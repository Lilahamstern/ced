package redis

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"os"
	"strconv"
	"time"
)

type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const aqpPrefix = "apq:"

func (c *Cache) Add(ctx context.Context, hash string, query interface{}) {
	c.client.Set(aqpPrefix+hash, query, c.ttl)
}

func (c *Cache) Get(ctx context.Context, hash string) (interface{}, bool) {
	s, err := c.client.Get(aqpPrefix + hash).Result()
	if err != nil {
		return "", false
	}
	return s, true
}

func NewCache() *Cache {
	db, _ := strconv.Atoi(os.Getenv("redisdb"))

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redisaddr"),
		Password: os.Getenv("redispass"),
		DB:       db,
	})

	err := client.Ping().Err()
	if err != nil {
		log.Fatal(err)
	}

	ttl := 24 * time.Hour

	log.Println("Connected to redis db...")
	return &Cache{client: client, ttl: ttl}
}
