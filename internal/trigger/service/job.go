package service

import (
	"context"
	"time"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/proto/task"
)

func newJob(spec string) *job {
	return &job{
		spec: spec,
	}
}

type job struct {
	spec string
}

func (j *job) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := taskService.DispatchTasks(ctx, &task.DispatchTasksRequest{Spec: j.spec})
	if err != nil {
		logger.Errorf("DispatchTasks: %s", err)
	}
	logger.Debugf("DispatchTasks-spec: %s", j.spec)
}
