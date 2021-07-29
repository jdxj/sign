package main

import (
	"fmt"
	"net/http"

	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/task"
	"github.com/jdxj/sign/internal/task/bili"
	"github.com/jdxj/sign/internal/task/common"
	"github.com/jdxj/sign/internal/task/hpi"
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
	var (
		err    error
		client *http.Client
	)

	for _, ud := range uds {
		switch ud.Domain {
		case common.BiliDomain:
			client, err = bili.Auth(ud.Key)

		case common.HPIDomain:
			client, err = hpi.Auth(ud.Key)

		default:
			err = fmt.Errorf("unsupport domain: %d", ud.Domain)
		}

		if err != nil {
			logger.Errorf("addVal err: %s, id: %s, domain: %s",
				err, ud.ID, ud.Domain)
			continue
		}

		for _, typ := range ud.Type {
			t := &common.Task{
				ID:     ud.ID,
				Type:   typ,
				Client: client,
			}

			switch typ {
			case common.BiliSign:
				bili.AddSignTask(t)

			case common.BiliBCount:
				bili.AddBCountTask(t)

			case common.HPISign:
				hpi.AddSignTask(t)

			default:
				err = fmt.Errorf("unsupport type: %d", typ)
			}

			if err != nil {
				logger.Warnf("addVal err: %s, id: %s, type: %d",
					err, ud.ID, typ)
				continue
			}
		}
	}
}
