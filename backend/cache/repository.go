package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type Repository interface {
	SetKey(name string, nameSpace string, value interface{}, expiration time.Duration) error
	GetKey(name string, nameSpace string, model interface{}) error
	DeleteKey(name string, nameSpace string) error
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
