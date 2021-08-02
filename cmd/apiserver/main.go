package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/storage"
	"github.com/jdxj/sign/internal/task"
)

func main() {
	flagSet := pflag.NewFlagSet("apiserver", pflag.ExitOnError)
	file := flagSet.StringP("file", "f", "config.yaml", "configure path")
	_ = flagSet.Parse(os.Args) // 忽略 err, 因为使用了 ExitOnError

	root := config.ReadConfigs(*file)
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))
	botCfg := root.Bot
	bot.Init(botCfg.Token, botCfg.ChatID)
	storage.Init(root.Storage.Path)

	task.RecoverTasks()
	task.Start()

	err := apiserver.Run(root.APIServer)
	if err != nil {
		logger.Errorf("api server run err: %s", err)
	}
}
