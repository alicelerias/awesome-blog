package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type Repository interface {
	SetKey(name string, id string, value interface{}, expiration time.Duration) error
	GetKey(name string, id string, model interface{}) error
	DeleteKey(name string, id string) error
}

type RedisClient struct {
	client *redis.Client
}

var redisClient Repository = &RedisClient{}

func NewRedisClient(client *redis.Client) *RedisClient {
	return &RedisClient{
		client: client,
	}
}
