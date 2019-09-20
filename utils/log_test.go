package utils

import (
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("test")
}

func TestMyLog(t *testing.T) {

}
