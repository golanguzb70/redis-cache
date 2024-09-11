package rediscache

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache interface {
	// Set stores a key-value pair in the cache.
	Set(ctx context.Context, key string, value string, expiration int) error
	// Get retrieves a key from the cache.
	Get(ctx context.Context, key string) (string, error)
	// Del deletes a key from the cache.
	Del(ctx context.Context, key string) error
	/*
		DelWildCard deletes all keys that match the wildcard pattern.
		Patterns examples:
		- "prefix:*"
		- "*:suffix"
		- "*" (deletes all keys)
	*/
	DelWildCard(ctx context.Context, wildcard string) error
	// Ping checks if the cache is available.
	Ping(ctx context.Context) error
	// HashOject generates a hash for the given object.
	HashOject(obj interface{}) string
	// Hash generates a hash for the given key.
	Hash(key string) string
}

type cache struct {
	client *redis.Client
}

func New(cfg *Config) (RedisCache, error) {
	options := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		DB:   0,
	}

	client := redis.NewClient(options)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		options.Username = cfg.RedisUsername
		options.Password = cfg.RedisPassword

		client = redis.NewClient(options)

		_, err = client.Ping(context.Background()).Result()
		if err != nil {
			return nil, err
		}
	}

	return &cache{
		client: client,
	}, nil
}

// Set stores a key-value pair in the cache.
func (c *cache) Set(ctx context.Context, key string, value string, expiration int) error {
	return c.client.Set(ctx, key, value, time.Duration(expiration)*time.Second).Err()
}

// Get retrieves a key from the cache.
func (c *cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// Del deletes a key from the cache.
func (c *cache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Ping checks if the cache is available.
func (c *cache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

/*
DelWildCard deletes all keys that match the wildcard pattern.
Patterns examples:
- "prefix:*"
- "*:suffix"
- "*" (deletes all keys)
*/
func (c *cache) DelWildCard(ctx context.Context, wildcard string) error {
	keys, err := c.client.Keys(ctx, wildcard).Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		if err := c.Del(ctx, key); err != nil {
			return err
		}
	}

	return nil
}

/*
HashOject generates a hash for the given object.
*/
func (c *cache) HashOject(obj interface{}) string {
	js, _ := json.Marshal(obj)
	hasher := sha256.New()

	hasher.Write([]byte(js))
	hashSum := hasher.Sum(nil)

	return hex.EncodeToString(hashSum)
}

/*
Hash generates a hash for the given key.
*/
func (c *cache) Hash(key string) string {
	hasher := sha256.New()

	hasher.Write([]byte(key))

	hashSum := hasher.Sum(nil)

	return hex.EncodeToString(hashSum)
}
