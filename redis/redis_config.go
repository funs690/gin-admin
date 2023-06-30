package redis

import (
	"context"
	"fmt"
	"gin-admin/config"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	Rdb *redis.Client
	ctx = context.Background()
)

// init redis client
func InitRedisClient() {
	// create database
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
}

// 设置数据, expiration = 0 keep stay
func Set(key string, data interface{}, expiration time.Duration) error {
	return Rdb.Set(ctx, key, data, expiration).Err()
}

// 获取数据
func Get(key string) (string, error) {
	return Rdb.Get(ctx, key).Result()
}

// 删除数据
func Delete(key string) error {
	return Rdb.Del(ctx, key).Err()
}
