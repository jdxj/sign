package task

import (
	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/pkg/bot"
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
	_, _ = c.AddFunc("29 19 * * *", func() {
		testNotify()
	})

	_, _ = c.AddFunc("30 19 * * *", bili.RunSignTask)
	_, _ = c.AddFunc("31 19 * * *", bili.RunBCountTask)

	_, _ = c.AddFunc("30 19 * * *", hpi.RunSignTask)
}

func testNotify() {
	text := `测试 apiserver, jenkins, retest`
	bot.Send(text)
}
