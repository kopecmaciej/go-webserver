package lib

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Redis struct {
	Client *redis.Client
}

type redisCache struct {
	host     string
	password string
	db       int
}

func InitRedis() (*Redis, error) {
  address := os.Getenv("REDIS_URL")
  password := os.Getenv("REDIS_PASSWORD")
  db := 0

  client := redis.NewClient(&redis.Options{
      Addr: address,
      Password: password,
      DB: db,
   })
   if err := client.Ping(ctx).Err(); err != nil {
      return nil, err
   }
   return &Redis{
      Client: client,
   }, nil
}

func (cache *redisCache) RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.password,
		DB:       cache.db,
	})
}

func (cache *redisCache) Get(key string) *redis.StringCmd {
	client := cache.RedisClient()

	return client.Get(ctx, key)
}

func (cache *redisCache) Set(key string, value interface{}, exp time.Duration) *redis.StatusCmd {
	client := cache.RedisClient()

	return client.Set(ctx, key, value, exp)
}
