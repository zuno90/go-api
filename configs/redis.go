package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
)

func ConnectRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:	  os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	pong, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong, "Redis is started")
	return redisClient
}