package task

import (
	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/pkg/task/bili"
	"github.com/jdxj/sign/pkg/task/hpi"
)

func Run() {
	c := cron.New()
	_, _ = c.AddFunc("0 20 * * *", bili.RunSignTask)
	_, _ = c.AddFunc("1 20 * * *", bili.RunBCountTask)

	_, _ = c.AddFunc("0 20 * * *", hpi.RunSignTask)
	c.Run()
}
