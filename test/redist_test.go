package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var RDB *gorm.DB

func TestRedis(t *testing.T) {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	rdb.Set("hello", "world", time.Minute*10)
	get := rdb.Get("hello")
	fmt.Println(get)
}
