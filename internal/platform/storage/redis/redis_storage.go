package redis

import (
	"context"
	"fmt"
	"github.com/alexperezortuno/go-auth/common/environment"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

// RedisService is struct wrapper around raw Redis client
type RedisService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &RedisService{}
	ctx          = context.Background()
)

var params = environment.Server()

const CacheDuration = 1 * time.Hour

// InitializeStore is initializing the store service and return a store pointer
func InitializeStore() *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     params.RedisHost,
		Password: params.RedisPass, // no password set
		DB:       params.RedisDb,   // use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(fmt.Sprintf("Error init Redis: %v", err))
	}

	log.Println(fmt.Sprintf("Redis started successfully: pong message = {%s}", pong))
	storeService.redisClient = rdb
	return storeService
}

func SaveURLInRedis(shortURL, originalURL string) {
	err := storeService.redisClient.Set(ctx, shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		log.Println(fmt.Sprintf("Failed SaveURLInRedis | Error: %v - shortURL: %s - originalURL: %s\n",
			err, shortURL, originalURL))
	}
}

func RetrieveInitialURLFromRedis(shortURL string) string {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		log.Println(fmt.Sprintf("Failed RetrieveInitialURLFromRedis | Error: %v - shortURL: %s\n",
			err, shortURL))
	}
	return result
}

func SaveTokenInRedis(token, username string) {
	err := storeService.redisClient.Set(ctx, token, username, CacheDuration).Err()
	if err != nil {
		log.Println(fmt.Sprintf("Failed SaveTokenInRedis | Error: %v - token: %s - username: %s\n",
			err, token, username))
	}
}
