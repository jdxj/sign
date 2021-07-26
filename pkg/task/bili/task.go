package bili

import (
	"github.com/jdxj/sign/pkg/task/common"
)

var (
	signTasks   = common.NewPool()
	bCountTasks = common.NewPool()
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

func AddBCountTask(task *common.Task) {
	bCountTasks.AddTask(task)
}

func RunBCountTask() {
	for num, task := range bCountTasks.GetAll() {
		success := false
		for i := 0; i < 3; i++ {
			if QueryBi(task) {
				success = true
				break
			}
		}

		if !success {
			bCountTasks.DelTask(num)
		}
	}
}
