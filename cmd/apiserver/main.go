package main

import (
	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
)

func main() {
	root := config.ReadConfigs("config.yaml")
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))

	err := apiserver.Run(root.APIServer)
	if err != nil {
		logger.Errorf("api server run err: %s", err)
	}
}
