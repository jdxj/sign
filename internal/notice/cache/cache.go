package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-micro.dev/v4/cache"

	"github.com/jdxj/sign/internal/pkg/logger"
)

const (
	cacheKeyUserTelegram = "USER_TELEGRAM"
)

var (
	cc         cache.Cache
	expiration = 5 * time.Minute
)

func init() {
	cc = cache.NewCache(cache.Expiration(expiration))
}

func getFullCacheKey(prefix string, suffix interface{}) string {
	return fmt.Sprintf("%s:%v", prefix, suffix)
}

func getUserTelegramCacheKey(userID int64) string {
	return getFullCacheKey(cacheKeyUserTelegram, userID)
}

func GetUserTelegram(ctx context.Context, userID int64) int64 {
	res, _, err := cc.Context(ctx).Get(getUserTelegramCacheKey(userID))
	if errors.Is(err, cache.ErrKeyNotFound) ||
		errors.Is(err, cache.ErrItemExpired) {
		return 0
	} else if err != nil {
		logger.Errorf("GetUserTelegram: %s, userID: %d", err, userID)
		return 0
	}
	return res.(int64)
}

func SetUserTelegram(ctx context.Context, userID, telegram int64) {
	err := cc.Context(ctx).Put(getUserTelegramCacheKey(userID), telegram, expiration)
	if err != nil {
		logger.Errorf("SetUserTelegram: %s, userID: %d, telegram: %d", err, userID, telegram)
	}
}
