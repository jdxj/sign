package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/jdxj/sign/internal/pkg/config"
)

var (
	Redis *redis.Client
)

func InitRedis(conf config.RDB) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Pass,
		DB:       conf.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := Redis.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
}
