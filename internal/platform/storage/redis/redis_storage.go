package redis

import (
	"context"
	"fmt"
	"github.com/alexperezortuno/go-auth/common"
	"github.com/alexperezortuno/go-auth/common/environment"
	"github.com/go-redis/redis/v9"
	"log"
)

// RedisService is struct wrapper around raw Redis client
type RedisService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService1 = &RedisService{}
	storeService2 = &RedisService{}
	ctx           = context.Background()
)

var env = environment.Server()

// InitializeStore is initializing the store service and return a store pointer
func InitializeStore() (*RedisService, *RedisService) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost,
		Password: env.RedisPass, // no password set
		DB:       env.RedisDb,   // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(fmt.Sprintf("Error init Redis: %v", err))
	}

	rdb2 := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost,
		Password: env.RedisPass, // no password set
		DB:       env.RedisDb2,  // use default DB
	})

	_, err2 := rdb2.Ping(ctx).Result()
	if err2 != nil {
		log.Println(fmt.Sprintf("Error init Redis: %v", err))
	}

	log.Println(fmt.Sprintf("Redis started successfully: pong message = {%s}", pong))

	storeService1.redisClient = rdb
	storeService2.redisClient = rdb2

	return storeService1, storeService2
}

func GetToken(token string) string {
	result, err := storeService1.redisClient.Get(ctx, token).Result()
	if err != nil {
		log.Println(fmt.Sprintf("Failed getting token | Error: %v - token: %s\n",
			err, token))
		return ""
	}
	return result
}

func SaveToken(token, username string) {
	err := storeService1.redisClient.Set(ctx, token, username, common.PrimaryCacheDuration()).Err()
	if err != nil {
		log.Println(fmt.Sprintf("Failed SaveToken | Error: %v - token: %s - username: %s\n",
			err, token, username))
	}
}

func DeleteToken(token string) {
	err := storeService1.redisClient.Del(ctx, token).Err()
	if err != nil {
		log.Println(fmt.Sprintf("Failed deleting token | Error: %v - token: %s\n",
			err, token))
	}
}

func SaveRefreshToken(refreshToken, username string) {
	err := storeService2.redisClient.Set(ctx, refreshToken, username, common.SecondaryCacheDuration()).Err()
	if err != nil {
		log.Println(fmt.Sprintf("Failed SaveRefreshToken | Error: %v - token: %s - username: %s\n",
			err, refreshToken, username))
	}
}

func GetRefreshToken(refreshToken string) string {
	result, err := storeService2.redisClient.Get(ctx, refreshToken).Result()
	if err != nil {
		log.Println(fmt.Sprintf("Failed getting token | Error: %v - token: %s\n",
			err, refreshToken))
		return ""
	}
	return result
}

func DeleteRefreshToken(refreshToken string) {
	err := storeService2.redisClient.Del(ctx, refreshToken).Err()
	if err != nil {
		log.Println(fmt.Sprintf("Failed deleting token | Error: %v - token: %s\n",
			err, refreshToken))
	}
}
