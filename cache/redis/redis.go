package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

// Config options.
type Config struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

// New  redis client.
func New(c Config) *redis.Client {
	if len(c.Addr) == 0 {
		panic("redis addr is nil")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
		PoolSize: c.PoolSize,
	})
	if err := rdb.Ping(context.Background()).Err(); nil != err {
		logx.Error(err)
		panic(err)
	}
	return rdb
}
