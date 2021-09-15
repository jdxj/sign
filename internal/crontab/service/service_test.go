package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func TestMain(t *testing.M) {
	logger.Init("./crontab.log")
	rpc.Init("127.0.0.1:2379")
	os.Exit(t.Run())
}

func TestService_CreateTask(t *testing.T) {
	var client crontab.CrontabServiceClient
	rpc.NewClient("crontab", func(cc *grpc.ClientConn) {
		client = crontab.NewCrontabServiceClient(cc)
	})

	rsp, err := client.CreateTask(context.Background(), &crontab.CreateTaskRep{
		UserID: 1,
		Task: &crontab.Task{
			Describe: "test create task 2",
			Kind:     102,
			Spec:     "* * * * *",
			SecretID: 1,
		},
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("taskID: %d\n", rsp.TaskID)
}
