package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr, port, password string, db int) (redis.UniversalClient, error) {
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{fmt.Sprintf("%s:%s", addr, port)},
		Password: password,
		DB:       db,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Could not connect to Redis: %v", err)
		return nil, err
	}

	log.Println("Connected to Redis successfully")
	return redisClient, nil
}
