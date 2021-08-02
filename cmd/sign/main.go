package main

import (
	"os"

	"github.com/spf13/pflag"

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
	bot.Init(root.Bot.Token, root.Bot.ChatID)
	storage.Init(root.Storage.Path)

	addVal(root.User)

	logger.Infof("started")
	task.RecoverTasks()
	task.Run()
}

func addVal(uds []config.User) {
	var err error
	for _, ud := range uds {
		t := &task.Task{
			ID:     ud.ID,
			Domain: ud.Domain,
			Types:  ud.Type,
			Key:    ud.Key,
		}
		err = task.Add(t)
		if err != nil {
			logger.Errorf("addVal err: %s, id: %s, domain: %s",
				err, ud.ID, ud.Domain)
			continue
		}
	}
}
