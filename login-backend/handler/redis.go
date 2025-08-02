package handler

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	Rdb *redis.Client
	Ctx context.Context
)

func SetRedisClient(r *redis.Client, c context.Context) {
	Rdb = r
	Ctx = c
}
