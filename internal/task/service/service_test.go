package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
)

var (
	client crontab.CrontabServiceClient
)

func TestMain(t *testing.M) {
	logger.Init("./crontab.log")
	rpc.Init("172.17.0.4:2379")
	rpc.NewClient("crontab", func(cc *grpc.ClientConn) {
		client = crontab.NewCrontabServiceClient(cc)
	})
	os.Exit(t.Run())
}

func TestService_CreateTask(t *testing.T) {
	rsp, err := client.CreateTask(context.Background(), &crontab.CreateTaskReq{
		UserID:   1,
		Describe: "test bili sign in",
		Kind:     crontab.Kind_BILISignIn,
		Spec:     "* * * * *",
		SecretID: 1,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("taskID: %d\n", rsp.TaskID)
}

func TestService_DeleteTask(t *testing.T) {
	_, err := client.DeleteTask(context.Background(), &crontab.DeleteTaskReq{
		TaskID: 5,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
