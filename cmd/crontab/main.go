package main

import (
	"os"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/crontab/service"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func main() {
	flagSet := pflag.NewFlagSet(crontab.ServiceName, pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	logger.Init(root.Logger.Path+crontab.ServiceName+".log",
		logger.WithMode(root.Logger.Mode))

	dbConf := root.DB
	db.InitGorm(dbConf)

	rpcConf := root.RPC
	server, err := rpc.NewServer(crontab.ServiceName,
		rpcConf.EtcdAddr, rpcConf.CrontabPort)
	if err != nil {
		logger.Errorf("new %s rpc server err: %s", crontab.ServiceName, err)
		return
	}

	srv := service.NewService()
	srv.Start()

	server.Register(func(registrar grpc.ServiceRegistrar) {
		crontab.RegisterCrontabServiceServer(registrar, srv)
	})
	server.Serve()

	util.Hold()

	server.Stop()
	srv.Stop()
}