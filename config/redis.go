package config

import (
	"fmt"
	"main/utils"

	"github.com/redis/go-redis/v9"
)

var RedisClient redis.Client

func ConnectRedis() {

  redisDb, err := utils.StringToNumber(EnvirontmentVariables.RedisDb)
  if(err != nil) {
    fmt.Printf("Error when try to connect to redis: %v\n", err)
    return;
  }

	client := redis.NewClient(&redis.Options{
		Addr:     EnvirontmentVariables.RedisHost,
		Password: EnvirontmentVariables.RedisPass,
		DB:       int(redisDb),
	})

  RedisClient = *client;

  fmt.Printf("Redis connected")
}
