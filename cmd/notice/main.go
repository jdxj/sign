package main

import (
	"os"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/notice/service"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/notice"
)

func main() {
	flagSet := pflag.NewFlagSet(notice.ServiceName, pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	logger.Init(root.Logger.Path+notice.ServiceName+".log",
		logger.WithMode(root.Logger.Mode))

	rpcConf := root.RPC
	rpc.Init(rpcConf.EtcdAddr)

	server, err := rpc.NewServer(notice.ServiceName,
		rpcConf.EtcdAddr, rpcConf.NoticePort)
	if err != nil {
		logger.Errorf("new %s rpc server err: %s", notice.ServiceName, err)
		return
	}

	botConf := root.Bot
	server.Register(func(registrar grpc.ServiceRegistrar) {
		notice.RegisterNoticeServiceServer(registrar, service.New(botConf))
	})
	server.Serve()
	util.Hold()
	server.Stop()
}
