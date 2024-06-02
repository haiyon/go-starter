package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-starter/pkg/log"

	"github.com/redis/go-redis/v9"
)

// ICache defines a general caching interface
type ICache[T any] interface {
	// Get data from cache using the specified field and return a pointer to type T and a possible error.
	Get(context.Context, string) (*T, error)
	// Set saves the specified data into cache using the specified string as key.
	Set(context.Context, string, *T) error
	// Delete data from cache using the specified key.
	Delete(context.Context, string) error
	// Reset data in cache using the specified pointer to type T as new value.
	Reset(context.Context, string, *T) error
}

// Cache implements the ICache interface
type Cache[T any] struct {
	rc      *redis.Client
	key     string
	useHash bool
}

// Key defines the cache key
func Key(key string) string {
	return fmt.Sprintf("%s", key)
}

// NewCache creates a new Cache instance
func NewCache[T any](rc *redis.Client, key string, useHash bool) *Cache[T] {
	return &Cache[T]{rc: rc, key: key, useHash: useHash}
}

// Get retrieves data from cache
func (c *Cache[T]) Get(ctx context.Context, field string) (*T, error) {
	var result string
	var err error

	if c.useHash {
		result, err = c.rc.HGet(ctx, c.key, field).Result()
	} else {
		result, err = c.rc.Get(ctx, field).Result()
	}

	if err != nil {
		// if errors.Is(err, redis.Nil) {
		// 	return nil, nil // Cache miss
		// }
		return nil, fmt.Errorf("failed to get cache: %w", err)
	}

	var row T
	if err = json.Unmarshal([]byte(result), &row); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cache data: %w", err)
	}
	return &row, nil
}

// Set saves data into cache
func (c *Cache[T]) Set(ctx context.Context, field string, data *T) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Errorf(nil, "failed to marshal data for cache set: %v, error: %v", data, err)
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if c.useHash {
		err = c.rc.HSet(ctx, c.key, field, bytes).Err()
	} else {
		err = c.rc.Set(ctx, field, bytes, 0).Err()
	}

	if err != nil {
		log.Errorf(nil, "failed to set cache: %v, error: %v", data, err)
		return fmt.Errorf("failed to set cache: %w", err)
	}
	return nil
}

// Delete removes data from cache
func (c *Cache[T]) Delete(ctx context.Context, field string) error {
	var err error

	if c.useHash {
		err = c.rc.HDel(ctx, c.key, field).Err()
	} else {
		err = c.rc.Del(ctx, field).Err()
	}

	if err != nil {
		log.Errorf(nil, "failed to delete cache field: %s, error: %v", field, err)
		return fmt.Errorf("failed to delete cache: %w", err)
	}
	return nil
}

// Reset resets data in cache
func (c *Cache[T]) Reset(ctx context.Context, field string, data *T) error {
	return c.Set(ctx, field, data)
}
