package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/jdxj/sign/internal/executor"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/mq"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/pkg/util"
)

const (
	serviceName = "executor"
)

func main() {
	flagSet := pflag.NewFlagSet(serviceName, pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	logger.Init(root.Logger.Path+serviceName+".log",
		logger.WithMode(root.Logger.Mode))

	rpcConf := root.RPC
	rpc.Init(rpcConf.EtcdAddr)

	rabbitConf := root.Rabbit
	mq.InitRabbit(rabbitConf)

	e := executor.New()
	e.Start()
	util.Hold()
	e.Stop()
}
