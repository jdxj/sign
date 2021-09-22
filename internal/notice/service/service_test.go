package service

import (
	"context"
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/proto/notice"
)

var (
	client notice.NoticeServiceClient
)

func TestMain(t *testing.M) {
	logger.Init("./notice.log")
	rpc.Init("127.0.0.1:2379")
	os.Exit(t.Run())
}

func TestService_SendMessage(t *testing.T) {
	rpc.NewClient(notice.ServiceName, func(cc *grpc.ClientConn) {
		client = notice.NewNoticeServiceClient(cc)
	})

	_, err := client.SendMessage(context.Background(), &notice.SendMessageReq{
		UserID: 1,
		Text:   "test notice service",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
