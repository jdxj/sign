package task

import (
	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/task/bili"
	"github.com/jdxj/sign/internal/task/hpi"
)

var (
	ic = cron.New()
)

func Run() {
	addTask(ic)
	ic.Run()
}

func Start() {
	addTask(ic)
	ic.Start()
}

func addTask(c *cron.Cron) {
	_, _ = c.AddFunc("0 20 * * *", bili.RunSignTask)
	_, _ = c.AddFunc("1 20 * * *", bili.RunBCountTask)

	_, _ = c.AddFunc("0 20 * * *", hpi.RunSignTask)
}
