package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"url-shortener/pkg/config"
)

// Client is the implementation of RedisClient
type Client struct {
	client     *redis.Client
	expiration time.Duration
}

// NewRedisClient create a new instance of the Redis client
func NewRedisClient(conf *config.Config) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.DatabasesConfig.Redis.Address,
		Password: conf.DatabasesConfig.Redis.Password,
		DB:       conf.DatabasesConfig.Redis.DB,
	})

	return &Client{
		client:     client,
		expiration: time.Duration(conf.DatabasesConfig.Redis.ExpirationMinutes) * time.Minute,
	}
}

// Ping check the connection to Redis
func (c *Client) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

// Set save a value in Redis with an expiration
func (c *Client) Set(ctx context.Context, key string, value interface{}) error {
	return c.client.Set(ctx, key, value, c.expiration).Err()
}

// Get search and return a value from Redis
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Client) Clean(ctx context.Context, key string) error {
	return c.client.Set(ctx, key, "", 5*time.Second).Err()
}

// Shutdown close the Redis client safely
func (c *Client) Shutdown() error {
	return c.client.Close()
}
