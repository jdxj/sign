package main

import (
	"github.com/jdxj/sign/internal/logger"
	"github.com/jdxj/sign/pkg/apiserver"
	"github.com/jdxj/sign/pkg/config"
)

func main() {
	root := config.ReadConfigs("config.yaml")
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))

	err := apiserver.Run(root.APIServer.Host, root.APIServer.Port)
	if err != nil {
		logger.Errorf("api server run err: %s", err)
	}
}
