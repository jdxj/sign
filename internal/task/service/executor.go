package service

import (
	"context"
	"time"

	"go-micro.dev/v4/broker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/proto/notice"
	pb "github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/task/service/executor"
)

type Executor interface {
	Kind() string
	Execute([]byte) string
}

var (
	executors = map[string]Executor{
		pb.Kind_MOCK.String(): &executor.MockExecutor{},
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
	content := exe.Execute(task.GetParam())
	// todo: 重试机制
	if err != nil {
		logger.Errorf("Execute, taskId: %d, err: %s", task.GetTaskId(), err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = noticeService.SendNotice(ctx, &notice.SendNoticeRequest{
		UserId:  task.GetUserId(),
		Content: content,
	})
	if err != nil {
		return status.Errorf(codes.Internal, "SendNotice: %s", err)
	}
	return nil
}
