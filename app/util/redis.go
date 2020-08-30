package util

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		// Password: "",
		DB: 0,
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		fmt.Println("redis ping error")
	}

	Redis.Set("", "hello world", 600)
}
