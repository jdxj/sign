package service

import (
	"os"
	"testing"

	"github.com/jdxj/sign/internal/pkg/logger"
)

func TestMain(t *testing.M) {
	logger.Init("./user.log")
	os.Exit(t.Run())
}
