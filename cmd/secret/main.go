package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/proto/secret"
	"github.com/jdxj/sign/internal/secret/service"
)

const (
	serviceName = "secret"
)

func main() {
	flagSet := pflag.NewFlagSet(serviceName, pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	logger.Init(root.Logger.Path+serviceName+".log",
		logger.WithMode(root.Logger.Mode))

	rpcConf := root.RPC
	server, err := rpc.NewServer(serviceName,
		rpcConf.EtcdAddr, rpcConf.SecretPort)
	if err != nil {
		logger.Errorf("new %s rpc server err: %s", serviceName, err)
		return
	}

	server.Register(func(registrar grpc.ServiceRegistrar) {
		secret.RegisterSecretServiceServer(registrar, &service.Service{})
	})
	server.Serve()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	logger.Infof("receive signal: %d", s)

	server.Stop()
}
