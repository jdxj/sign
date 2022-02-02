package service

import (
	"go-micro.dev/v4/client"

	"github.com/jdxj/sign/internal/proto/task"
)

var (
	taskService task.TaskService
)

func Init(cc client.Client) error {
	taskService = task.NewTaskService(task.ServiceName, cc)
	return nil
}
