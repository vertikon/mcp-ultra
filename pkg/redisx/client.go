package redisx

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Client is a facade for Redis client operations
type Client struct {
	inner *redis.Client
}

// NewClient wraps a redis.Client
func NewClient(client *redis.Client) *Client {
	return &Client{inner: client}
}

// Inner returns the underlying redis.Client for advanced operations
func (c *Client) Inner() *redis.Client {
	return c.inner
}

// Cmdable returns the client as redis.Cmdable interface for compatibility
func (c *Client) Cmdable() redis.Cmdable {
	return c.inner
}

// Set stores a value with expiration
func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.inner.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	result, err := c.inner.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", ErrKeyNotFound
	}
	return result, err
}

// Del deletes one or more keys
func (c *Client) Del(ctx context.Context, keys ...string) error {
	return c.inner.Del(ctx, keys...).Err()
}

// Exists checks if a key exists (returns bool for single key)
func (c *Client) Exists(ctx context.Context, key string) (bool, error) {
	n, err := c.inner.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// Incr increments a counter
func (c *Client) Incr(ctx context.Context, key string) (int64, error) {
	return c.inner.Incr(ctx, key).Result()
}

// SetNX sets a value only if it doesn't exist
func (c *Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.inner.SetNX(ctx, key, value, expiration).Result()
}

// SetEx sets a value with expiration
func (c *Client) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.inner.SetEx(ctx, key, value, expiration).Err()
}

// TTL returns time-to-live for a key
func (c *Client) TTL(ctx context.Context, key string) (time.Duration, error) {
	return c.inner.TTL(ctx, key).Result()
}

// FlushAll removes all keys
func (c *Client) FlushAll(ctx context.Context) error {
	return c.inner.FlushAll(ctx).Err()
}

// Ping tests the connection
func (c *Client) Ping(ctx context.Context) error {
	return c.inner.Ping(ctx).Err()
}
