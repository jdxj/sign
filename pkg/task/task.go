package task

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/internal/logger"
	"github.com/jdxj/sign/pkg/task/bili"
)

const (
	unknownTask = iota
	BiliSign
	BiliBCount
)

var (
	typeMap = map[int]string{
		BiliSign: "b站签到",
	}
)

var (
	tplSuccess = `%s 的 %s 任务执行成功`
	tplFailed  = `%s 的 %s 任务执行失败`
)

var (
	num   int
	tasks = make(map[int]*Task)
)

func AddTask(t *Task) {
	num++
	tasks[num] = t
}

func DelTask(num int) {
	delete(tasks, num)
}

type Task struct {
	ID     string
	Type   int
	Client *http.Client
}

func Run() {
	c := cron.New()
	_, err := c.AddFunc("0 8 * * *", bSignTask)
	if err != nil {
		log.Fatalln(err)
	}
	c.Run()
}

func bSignTask() {
	var (
		err  error
		text string
	)
	for num, task := range tasks {
		switch task.Type {
		case BiliSign:
			err = bili.SignIn(task.Client)
		default:
			logger.Warnf("unsupported task type: %d", task.Type)
			continue
		}

		if err == nil {
			text = fmt.Sprintf(tplSuccess, task.ID, typeMap[task.Type])
		} else {
			text = fmt.Sprintf(tplFailed+", err: %s",
				task.ID, typeMap[task.Type], err)
			DelTask(num)
		}

		err = bot.Send(text)
		if err != nil {
			logger.Errorf("send err: %s", err)
		}
	}
}
