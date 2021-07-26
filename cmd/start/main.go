package main

import (
	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/internal/logger"
	"github.com/jdxj/sign/pkg/config"
	"github.com/jdxj/sign/pkg/task"
)

func main() {
	root := config.ReadConfigs("config.yaml")

	bot.Init(root.Bot.Token, root.Bot.ChatID)
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))

	logger.Infof("started")
	task.Run()
}
