package main

import (
	"fmt"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/internal/logger"
	"github.com/jdxj/sign/pkg/config"
	"github.com/jdxj/sign/pkg/task"
	"github.com/jdxj/sign/pkg/task/bili"
	"github.com/jdxj/sign/pkg/task/common"
)

func main() {
	root := config.ReadConfigs("config.yaml")
	logger.Init(root.Logger.Path, logger.WithMode(root.Logger.Mode))
	bot.Init(root.Bot.Token, root.Bot.ChatID)

	addVal(root.User)

	logger.Infof("started")
	task.Run()
}

// todo: 使用 api
func addVal(uds []config.User) {
	var err error

	for _, ud := range uds {
		switch ud.Type {
		case common.BiliSign:
			err = bili.AddSignTask(ud.ID, ud.Key, ud.Type)
		default:
			err = fmt.Errorf("unsupport type: %d", ud.Type)
		}

		if err != nil {
			logger.Errorf("addVal err: %s, id: %s, type: %s",
				err, ud.ID, common.TypeMap[ud.Type])
			continue
		}
	}
}
