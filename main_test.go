package rediscache

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Define a test object struct
type TestObject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TestRedisCache is a struct to hold the redis cache instance for testing
type TestRedisCache struct {
	RedisCache
}

// NewTestRedisCache creates a new RedisCache instance for testing
func NewTestRedisCache(cfg *Config) *TestRedisCache {
	return &TestRedisCache{
		RedisCache: New(cfg),
	}
}

// TestSetGetDelete tests the Set, Get, and Delete methods of RedisCache
func (trc *TestRedisCache) TestSetGetDelete(t *testing.T) {
	ctx := context.Background()

	key := "test_key"
	value := "test_value"

	// Set value
	err := trc.Set(ctx, key, value, 10)
	assert.NoError(t, err, "Error setting value to cache")

	time.Sleep(2 * time.Second)

	// Get value
	result, err := trc.Get(ctx, key)
	assert.NoError(t, err, "Error getting value from cache")
	assert.Equal(t, value, result, "Retrieved value does not match the original value")

	key = "test_key2"
	value = "test_value2"
	err = trc.Set(ctx, key, value, 1)
	assert.NoError(t, err, "Error setting value to cache")

	// Wait for value to expire
	time.Sleep(2 * time.Second)

	// Verify value is expired
	_, err = trc.Get(ctx, key)
	assert.Error(t, err, "Value should have been expired from cache")

	err = trc.Set(ctx, key, value, 10)
	assert.NoError(t, err, "Error setting value to cache")

	// Delete value
	err = trc.Del(ctx, key)
	assert.NoError(t, err, "Error deleting value from cache")

	// Verify value is deleted
	_, err = trc.Get(ctx, key)
	assert.Error(t, err, "Value should have been deleted from cache")
}

// TestHash tests the Hash method of RedisCache
func (trc *TestRedisCache) TestHash(t *testing.T) {
	key := "test_key"
	expectedHash := "92488e1e3eeecdf99f3ed2ce59233efb4b4fb612d5655c0ce9ea52b5a502e655"

	hash := trc.Hash(key)
	assert.Equal(t, expectedHash, hash, "Hash does not match expected value")
}

// TestHashObject tests the HashObject method of RedisCache
func (trc *TestRedisCache) TestHashObject(t *testing.T) {
	testObj := TestObject{ID: 123, Name: "test"}

	expectedHash := "f99d2e506f11e31bc9102307e3ab14e170fbce1dbd244fb2d38f9e5753fbb397"

	hash := trc.HashOject(testObj)
	assert.Equal(t, expectedHash, hash, "Hash does not match expected value")
}

func TestCache(m *testing.T) {
	// Setup test configuration
	cfg := &Config{
		RedisHost:     "localhost",
		RedisPort:     "6379",
		RedisPassword: "", // No password for testing
	}
	// Create RedisCache instance for testing
	trc := NewTestRedisCache(cfg)

	// Run tests
	trc.TestSetGetDelete(m)
	trc.TestHash(m)
	trc.TestHashObject(m)

}
