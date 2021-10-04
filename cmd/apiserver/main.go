package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/apiserver/router"
	"github.com/jdxj/sign/internal/pkg/api"
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
	logger.Init(root.Logger.Path+"apiserver.log", logger.WithMode(root.Logger.Mode))

	rpcConf := root.RPC
	rpc.Init(rpcConf.EtcdAddr)

	apiConf := root.APIServer
	apiserver.Init(apiConf)

	r := router.New()
	server := api.NewServer(apiConf.Host, apiConf.Port, r)
	server.Start()
	util.Hold()
	server.Stop()
}
