package main

import (
	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/task"
)

func main() {
	root := config.ReadConfigs("config.yaml")
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))
	bot.Init(root.Bot.Token, root.Bot.ChatID)

	addVal(root.User)

	logger.Infof("started")
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
