package task

import (
	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/logger"
	"github.com/jdxj/sign/pkg/task/bili"
)

func Run() {
	c := cron.New()
	//spec := "* * * * *"
	//spec := "0 8 * * *"
	_, err := c.AddFunc("0 8 * * *", bili.RunSignTask)
	if err != nil {
		logger.Errorf("add func: %s", err)
		return
	}

	_, err = c.AddFunc("0 20 * * *", bili.RunBCountTask)
	if err != nil {
		logger.Errorf("add func: %s", err)
		return
	}

	c.Run()
}
