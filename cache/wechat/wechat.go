package wechat

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Wechat .
type Wechat struct {
	RDB *redis.Client
}

// New Wechat cache.
func New(rdb *redis.Client) *Wechat {
	return &Wechat{
		RDB: rdb,
	}
}

// Get .
func (wx *Wechat) Get(key string) interface{} {
	result, err := wx.RDB.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println(err)
	}
	if result == "" {
		return nil
	}
	return result
}

// Set .
func (wx *Wechat) Set(key string, val interface{}, expiration time.Duration) error {
	err := wx.RDB.Set(context.Background(), key, val, expiration).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// IsExist .
func (wx *Wechat) IsExist(key string) bool {
	var isExist bool
	if wx.RDB.Exists(context.Background(), key).Val() > 0 {
		isExist = true
	}
	return isExist
}

// Delete .
func (wx *Wechat) Delete(key string) error {
	err := wx.RDB.Del(context.Background(), key).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
