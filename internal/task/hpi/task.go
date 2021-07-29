package hpi

import "github.com/jdxj/sign/internal/task/common"

var (
	signTasks = common.NewPool()
)

func AddSignTask(task *common.Task) {
	signTasks.AddTask(task)
}

func RunSignTask() {
	for num, task := range signTasks.GetAll() {
		success := false
		for i := 0; i < 3; i++ {
			if SignIn(task) {
				success = true
				break
			}
		}

		if !success {
			signTasks.DelTask(num)
		}
	}
}
