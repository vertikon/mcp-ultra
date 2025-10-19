package redisx

import "github.com/redis/go-redis/v9"

// Options represents Redis connection options
type Options struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

// ToRedisOptions converts to redis.Options
func (o *Options) ToRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     o.Addr,
		Password: o.Password,
		DB:       o.DB,
		PoolSize: o.PoolSize,
	}
}

// NewClientFromOptions creates a new Client from options
func NewClientFromOptions(opts *Options) *Client {
	redisClient := redis.NewClient(opts.ToRedisOptions())
	return NewClient(redisClient)
}
