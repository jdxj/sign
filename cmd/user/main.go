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
	"github.com/jdxj/sign/internal/proto/user"
	"github.com/jdxj/sign/internal/user/service"
)

func main() {
	flagSet := pflag.NewFlagSet(user.ServiceName, pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	loggerConf := root.Logger
	logger.Init(loggerConf.Path+user.ServiceName+".log",
		logger.WithMode(loggerConf.Mode))

	dbConf := root.DB
	db.InitGorm(dbConf)

	rpcConf := root.RPC
	server, err := rpc.NewServer(user.ServiceName,
		rpcConf.EtcdAddr, rpcConf.UserPort, rpc.WithMod(loggerConf.Mode))
	if err != nil {
		logger.Errorf("new %s rpc server err: %s", user.ServiceName, err)
		return
	}

	server.Register(func(registrar grpc.ServiceRegistrar) {
		user.RegisterUserServiceServer(registrar, &service.Service{})
	})
	server.Serve()
	util.Hold()
	server.Stop()
}
