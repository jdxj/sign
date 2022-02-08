package cache

import (
	"context"
	"testing"
)

func TestGetUserTelegram(t *testing.T) {
	var (
		userID   int64 = 123
		telegram int64 = 456
	)
	SetUserTelegram(context.Background(), userID, telegram)

	res := GetUserTelegram(context.Background(), userID)
	if res != telegram {
		t.Fatalf("set telegram err")
	}
}
