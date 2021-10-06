package main

import (
	"os"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/secret"
	"github.com/jdxj/sign/internal/secret/service"
)

func main() {
	flagSet := pflag.NewFlagSet(secret.ServiceName, pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	loggerConf := root.Logger
	logger.Init(loggerConf.Path+secret.ServiceName+".log",
		logger.WithMode(loggerConf.Mode))

	dbConf := root.DB
	db.InitGorm(dbConf)

	rpcConf := root.RPC
	server, err := rpc.NewServer(secret.ServiceName,
		rpcConf.EtcdAddr, rpcConf.SecretPort, rpc.WithMod(loggerConf.Mode))
	if err != nil {
		logger.Errorf("new %s rpc server err: %s", secret.ServiceName, err)
		return
	}

	server.Register(func(registrar grpc.ServiceRegistrar) {
		secret.RegisterSecretServiceServer(registrar, service.New(root.Secret))
	})
	server.Serve()
	util.Hold()
	server.Stop()
}
