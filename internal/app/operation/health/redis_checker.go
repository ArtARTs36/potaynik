package health

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisChecker struct {
	redis *redis.Client
}

func NewRedisChecker(redis *redis.Client) *RedisChecker {
	return &RedisChecker{redis: redis}
}

func (c *RedisChecker) Check(ctx context.Context) CheckResult {
	return CheckResult{
		Service: "redis",
		Status:  c.redis.Ping(ctx).Err() == nil,
	}
}
