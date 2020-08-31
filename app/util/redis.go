package util

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		// Password: "",
		DB: 0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis ping error")
	} else {
		fmt.Println("redis connect success")
	}
}
