package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/apiserver/comm"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/pkg/util"
)

func main() {
	flagSet := pflag.NewFlagSet("apiserver", pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))

	rpcConf := root.RPC
	rpc.Init(rpcConf.EtcdAddr)
	comm.Init(root.APIServer)

	apiserver.Start(root.APIServer)
	util.Hold()
	apiserver.Stop()
}
