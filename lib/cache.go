package lib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient *redis.Client

type Redis struct {
	Client *redis.Client
}

type RedisCache struct {
	Redis
}

func InitRedis() {
	address := os.Getenv("REDIS_URL")
	password := os.Getenv("REDIS_PASSWORD")
	db := 0
	if len(address) == 0 {
		address = "localhost:6379"
	}
	if len(password) == 0 {
		password = "password"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
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
