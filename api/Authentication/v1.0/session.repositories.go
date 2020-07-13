package authentication

import (
	redisClient "Gogin/internal/platform/redis"
	"context"
	"os"
	"strconv"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// CreateSession Repository
func (repository SessionRepository) CreateSession(s Session) error {
	// Cleanup expired sessions
	redisClient.AuthClient.ZRemRangeByScore(ctx,
		"Auth:User:"+strconv.Itoa(s.UserID),
		"-inf",
		strconv.FormatInt(time.Now().Unix(), 10),
	)

	// Create new session with expiry
	duration, _ := strconv.Atoi(os.Getenv("AUTH_EXP"))
	return redisClient.AuthClient.ZAdd(ctx,
		"Auth:User:"+strconv.Itoa(s.UserID),
		&redis.Z{Score: float64(s.Timestamp + int64(duration)), Member: s.Token},
	).Err()
}

// DeleteSession Repository
func (repository SessionRepository) DeleteSession(s Session) {
	redisClient.AuthClient.ZRem(ctx, "Auth:User:"+strconv.Itoa(s.UserID), s.Token)
}

// IsSessionExisted Repository
func (repository SessionRepository) IsSessionExisted(s Session) bool {
	if redisClient.AuthClient.ZRank(ctx, "Auth:User:"+strconv.Itoa(s.UserID), s.Token).Err() != nil {
		return false
	}
	return true
}
