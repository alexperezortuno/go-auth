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
	environ := getEnvStr("APP_ENV", "prod")

	switch environ {
	case "dev":
		gin.SetMode(gin.DebugMode)
		break
	case "test":
		gin.SetMode(gin.TestMode)
		break
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		break
	default:
		log.Fatalf("invalid environment: %s", environ)
	}
}

func getEnvStr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

func Server() ServerValues {
	env()
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8082
	}

	protocol := getEnvStr("APP_PROTOCOL", "http")
	host := getEnvStr("APP_HOST", "localhost")
	timeZone := getEnvStr("APP_TIME_ZONE", "UTC")
	context := getEnvStr("APP_CONTEXT", "auth")
	redisHost := getEnvStr("REDIS_HOST", "")
	redisPass := getEnvStr("REDIS_PASS", "")
	redisDb := getEnvInt("REDIS_DB", 0)
	redisDb2 := getEnvInt("REDIS_DB_SECONDARY", 1)
	redisPort := getEnvInt("REDIS_PORT", 0)
	dbHost := getEnvStr("DB_HOST", "db-postgresql")
	dbUser := getEnvStr("DB_USER", "postgres")
	dbPass := getEnvStr("DB_PASS", "Me.123")
	dbPort := getEnvStr("DB_PORT", "5432")
	dbName := getEnvStr("DB_NAME", "authdb")
	dbTimeZone := getEnvStr("DB_TIME_ZONE", "UTC")
	engineSql := getEnvStr("DB_DRIVER", "postgres")
	tokenLifeTime := getEnvInt("TOKEN_LIFE_TIME", 15)
	refreshTokenLifeTime := getEnvInt("REFRESH_TOKEN_LIFE_TIME", 1)

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
