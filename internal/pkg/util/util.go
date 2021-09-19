package util

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jdxj/sign/internal/pkg/logger"
)

func Hold() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	logger.Infof("receive signal: %d", s)
}
