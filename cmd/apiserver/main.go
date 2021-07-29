package main

import (
	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/task"
)

func main() {
	root := config.ReadConfigs("config.yaml")
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))
	botCfg := root.Bot
	bot.Init(botCfg.Token, botCfg.ChatID)

	task.Start()

	err := apiserver.Run(root.APIServer)
	if err != nil {
		logger.Errorf("api server run err: %s", err)
	}
}
