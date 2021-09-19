package cache

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jdxj/sign/internal/pkg/config"
)

func TestMain(t *testing.M) {
	conf := config.RDB{
		Host: "127.0.0.1",
		Port: 6379,
		Pass: "",
		DB:   0,
	}
	InitRedis(conf)
	os.Exit(t.Run())
}

func TestRedis(t *testing.T) {
	cmd := Redis.Get(context.Background(), "hello")
	fmt.Printf("value: %s\n", cmd.Val())
}
