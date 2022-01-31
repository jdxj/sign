package service

import (
	"context"
	"time"

	"go-micro.dev/v4/broker"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/proto/notice"
	pb "github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/task/service/executor"
	"github.com/jdxj/sign/internal/task/service/executor/bilibili"
	"github.com/jdxj/sign/internal/task/service/executor/evo"
)

type Executor interface {
	Kind() string
	Execute([]byte) (string, error)
}

var (
	executors = map[string]Executor{
		pb.Kind_MOCK.String():              &executor.MockExecutor{},
		pb.Kind_BILIBILI_SIGN_IN.String():  &bilibili.SignIn{},
		pb.Kind_BILIBILI_B_COUNT.String():  &bilibili.Bi{},
		pb.Kind_EVOLUTION_RAPHAEL.String(): &evo.Updater{},
	}
)

func execute(e broker.Event) error {
	task := &pb.Task{}
	err := proto.Unmarshal(e.Message().Body, task)
	if err != nil {
		logger.Errorf("Unmarshal: %s", err)
		return err
	}

	exe, ok := executors[task.GetKind()]
	if !ok {
		logger.Errorf("kind not found, taskId: %d, kind: %s", task.GetTaskId(), task.GetKind())
		return nil
	}
	tryExecute(task, exe)
	return nil
}

func tryExecute(task *pb.Task, exe Executor) {
	var (
		retry    = 3
		interval = 3 * time.Second

		text string
		err  error
	)

	for i := 0; i < retry; i++ {
		text, err = exe.Execute(task.GetParam())
		if err != nil {
			logger.Errorf("try execute failed: %d, userID: %d, taskID: %d, err: %s",
				i, task.GetUserId(), task.GetTaskId(), err)
		} else {
			break
		}
		if i != retry-1 {
			time.Sleep(interval)
		}
	}
	if text != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = noticeService.SendNotice(ctx, &notice.SendNoticeRequest{
			UserId:  task.GetUserId(),
			Content: text,
		})
		if err != nil {
			logger.Errorf("SendNotice: %s", err)
		}
	}
}
