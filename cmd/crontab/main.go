package main

import (
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/crontab/service"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/proto/crontab"
)

const (
	serviceName = "crontab"
)

func main() {
	logger.Init("./crontab.log")

	// todo: 配置化地址
	server, err := rpc.NewServer(serviceName,
		"127.0.0.1:49152", "127.0.0.1:2379")
	if err != nil {
		logger.Errorf("new server err: %s", err)
		return
	}

	server.Register(func(registrar grpc.ServiceRegistrar) {
		crontab.RegisterTestServiceServer(registrar, &service.Service{})
	})
	server.Serve()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	logger.Infof("receive signal: %d", s)

	server.Stop()
}
