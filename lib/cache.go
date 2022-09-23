package lib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-web-server/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
	address     = &config.GlobalConfig.Cache.Url
	password    = &config.GlobalConfig.Cache.Password
)

type Redis struct {
	Client *redis.Client
}

type RedisCache struct {
	Redis
}

func InitRedis() {
	if len(*password) == 0 {
		*password = "password"
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     *address,
		Password: *password,
		DB:       0,
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	fmt.Println("redis connected")
}

func (cache *RedisCache) RedisClient() *redis.Client {
	return redisClient
}

func (cache *RedisCache) Get(key string) (*redis.StringCmd, error) {
	client := cache.RedisClient()
	val := client.Get(ctx, key)
	if val == nil {
		return nil, errors.New("record not found")
	}
	return val, nil
}

func (cache *RedisCache) Set(key string, value interface{}, exp time.Duration) *redis.StatusCmd {
	json, _ := json.Marshal(value)
	client := cache.RedisClient()
	return client.Set(ctx, key, json, exp)
}
