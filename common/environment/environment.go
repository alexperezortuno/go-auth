package environment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"time"
)

type ServerValues struct {
	Protocol             string
	Host                 string
	Port                 int
	ShutdownTimeout      time.Duration
	Context              string
	TimeZone             string
	RedisHost            string
	RedisPass            string
	RedisPort            int
	RedisDb              int
	RedisDb2             int
	DbUser               string
	DbPass               string
	DbHost               string
	DbPort               string
	DbName               string
	DbTimeout            time.Duration
	DbTimeZone           string
	EngineSql            string
	TokenLifeTime        int
	RefreshTokenLifeTime int
}

func env() {
	env := os.Getenv("APP_ENV")

	if env == "" || env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func getEnv(envName, valueDefault string) string {
	value := os.Getenv(envName)
	if value == "" {
		return valueDefault
	}
	return value
}

func getEnvInt(envName string, valueDefault int) int {
	value, err := strconv.Atoi(envName)
	if err != nil {
		return valueDefault
	}
	return value
}

func Server() ServerValues {
	env()
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8082
	}

	protocol := os.Getenv("APP_PROTOCOL")
	host := os.Getenv("APP_HOST")
	timeZone := os.Getenv("APP_TIME_ZONE")
	context := os.Getenv("APP_CONTEXT")
	redisHost := getEnv("REDIS_HOST", "")
	redisPass := getEnv("REDIS_PASS", "")
	redisDb := getEnvInt("REDIS_DB", 0)
	redisDb2 := getEnvInt("REDIS_DB_SECONDARY", 1)
	redisPort := getEnvInt("REDIS_DB", 0)
	dbHost := getEnv("DB_HOST", "db-postgresql")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASS", "Me.123")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "authdb")
	dbTimeZone := getEnv("DB_TIME_ZONE", "America/Santiago")
	engineSql := getEnv("DB_DRIVER", "postgres")
	tokenLifeTime := getEnvInt("TOKEN_LIFE_TIME", 15)
	refreshTokenLifeTime := getEnvInt("REFRESH_TOKEN_LIFE_TIME", 1)

	if err != nil {
		redisDb = 0
	}

	if protocol == "" {
		protocol = "http"
	}

	if host == "" {
		host = "localhost"
	}

	if context == "" {
		context = "auth"
	}

	if timeZone == "" {
		timeZone = "America/Santiago"
	}

	if redisPort == 0 {
		redisPort = 6379
	}

	if redisHost == "" {
		redisHost = fmt.Sprintf("redis:%d", redisPort)
	}

	log.Println(fmt.Printf("Redis host: %s, Redis pass: %s, Redis db: %s", redisHost, redisPass, redisDb))

	return ServerValues{
		Protocol:             protocol,
		Host:                 host,
		Context:              context,
		Port:                 port,
		TimeZone:             timeZone,
		ShutdownTimeout:      10 * time.Second,
		RedisHost:            redisHost,
		RedisPass:            redisPass,
		RedisDb:              redisDb,
		RedisDb2:             redisDb2,
		RedisPort:            redisPort,
		DbHost:               dbHost,
		DbPort:               dbPort,
		DbUser:               dbUser,
		DbPass:               dbPass,
		DbName:               dbName,
		DbTimeZone:           dbTimeZone,
		EngineSql:            engineSql,
		TokenLifeTime:        tokenLifeTime,
		RefreshTokenLifeTime: refreshTokenLifeTime,
	}
}
