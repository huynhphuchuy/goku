package redis

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// AuthClient redis
var AuthClient *redis.Client

func init() {
	authDB, _ := strconv.Atoi(os.Getenv("REDIS_AUTH_DB"))
	AuthClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       authDB,
	})
}
