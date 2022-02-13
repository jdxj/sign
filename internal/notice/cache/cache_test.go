package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserTelegram(t *testing.T) {
	var (
		userID   int64 = 123
		telegram int64 = 456
	)
	SetUserTelegram(context.Background(), userID, telegram)
	res := GetUserTelegram(context.Background(), userID)
	assert.Equal(t, telegram, res, "should equal")
}
