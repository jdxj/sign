package service

import (
	"os"
	"testing"

	"github.com/jdxj/sign/internal/pkg/logger"
)

func TestMain(t *testing.M) {
	logger.Init("./crontab.log")
	os.Exit(t.Run())
}
