package bili

import (
	"github.com/jdxj/sign/pkg/task/common"
)

var (
	signTasks = common.NewPool()
)

func AddSignTask(id, cookies string, typ int) error {
	client, err := Auth(cookies)
	if err != nil {
		return err
	}
	task := &common.Task{
		ID:     id,
		Type:   typ,
		Client: client,
	}
	signTasks.AddTask(task)
	return nil
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
