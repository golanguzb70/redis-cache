package rediscache

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache interface {
	Set(ctx context.Context, key string, value string, expiration int) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	Ping(ctx context.Context) error
	HashOject(obj interface{}) string
	Hash(key string) string
}

type cache struct {
	client *redis.Client
}

func New(cfg *Config) RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	return &cache{
		client: client,
	}
}

func (c *cache) Set(ctx context.Context, key string, value string, expiration int) error {
	return c.client.Set(ctx, key, value, time.Duration(expiration) * time.Second).Err()
}

func (c *cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *cache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *cache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

func (c *cache) HashOject(obj interface{}) string {
	js, _ := json.Marshal(obj)
	hasher := sha256.New()

	hasher.Write([]byte(js))
	hashSum := hasher.Sum(nil)

	return hex.EncodeToString(hashSum)
}

func (c *cache) Hash(key string) string {
	hasher := sha256.New()

	hasher.Write([]byte(key))

	hashSum := hasher.Sum(nil)

	return hex.EncodeToString(hashSum)
}
