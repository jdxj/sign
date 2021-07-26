package main

import (
	"fmt"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/internal/logger"
	"github.com/jdxj/sign/pkg/config"
	"github.com/jdxj/sign/pkg/storage"
	"github.com/jdxj/sign/pkg/task"
	"github.com/jdxj/sign/pkg/toucher"
)

func main() {
	root := config.ReadConfigs("config.yaml")

	bot.Init(root.Bot.Token, root.Bot.ChatID)
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))

	addVal(root.User)

	logger.Infof("started")
	task.Run()
}

// todo: 使用 api
func addVal(uds []config.User) {
	var (
		val toucher.Validator
		err error
		ds  = storage.Default
	)

	for _, ud := range uds {
		switch ud.Domain {
		case toucher.DomainBili:
			val, err = toucher.NewBili(ud.ID, ud.Key)
		default:
			err = fmt.Errorf("%w: %s",
				toucher.ErrorUnsupportedDomain, ud.Domain)
		}

		if err != nil {
			logger.Errorf("addVal err: %s", err)
			continue
		}
		ds.AddUserData(val)
	}
}
