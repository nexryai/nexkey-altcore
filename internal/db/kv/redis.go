package kv

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
)

func ConnectToRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Pass,
		//DB:       config.Redis.DB,
	})

	return client
}
