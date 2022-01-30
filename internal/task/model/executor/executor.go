package executor

import (
	"context"
	"fmt"
	"time"

	"go-micro.dev/v4/broker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/jdxj/sign/internal/proto/notice"
	pb "github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/task/client"
)

type Executor interface {
	Kind() string
	Execute([]byte) (string, error)
}

var (
	executors = map[string]Executor{
		pb.Kind_MOCK.String(): &mockExecutor{},
	}
)

func Execute(e broker.Event) error {
	task := &pb.Task{}
	err := proto.Unmarshal(e.Message().Body, task)
	if err != nil {
		return err
	}

	exe, ok := executors[task.GetKind()]
	if !ok {
		return fmt.Errorf("kind not found: %d", task.GetTaskId())
	}
	content, err := exe.Execute(task.GetParam())
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.NoticeService.SendNotice(ctx, &notice.SendNoticeRequest{
		UserId:  task.GetUserId(),
		Content: content,
	})
	if err != nil {
		return status.Errorf(codes.Internal, "SendNotice: %s", err)
	}
	return nil
}
