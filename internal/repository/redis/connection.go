package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/vertikon/mcp-ultra/internal/config"
	"github.com/vertikon/mcp-ultra/pkg/redisx"
)

// NewClient creates a new Redis client
func NewClient(cfg config.RedisConfig) *redisx.Client {
	client := redisx.NewClientFromOptions(&redisx.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	return client
}

// Ping tests Redis connection with timeout
func Ping(client *redisx.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := client.Ping(ctx)
	if err != nil {
		return fmt.Errorf("pinging Redis: %w", err)
	}
	return nil
}
