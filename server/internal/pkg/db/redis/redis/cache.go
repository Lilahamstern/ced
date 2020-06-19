package redis

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const aqpPrefix = "apq:"

func NewCache() *Cache {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := client.Ping().Err()
	if err != nil {
		log.Fatal(err)
	}

	ttl := 24 * time.Hour

	return &Cache{client: client, ttl: ttl}
}

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
