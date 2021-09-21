package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/proto/user"
)

var (
	client user.UserServiceClient
)

func TestMain(t *testing.M) {
	logger.Init("./user.log")
	rpc.Init("127.0.0.1:2379")
	rpc.NewClient(user.ServiceName, func(cc *grpc.ClientConn) {
		client = user.NewUserServiceClient(cc)
	})
	os.Exit(t.Run())
}

func TestService_CreateUser(t *testing.T) {
	rsp, err := client.CreateUser(context.Background(), &user.CreateUserReq{
		Nickname: "jdxj",
		Password: "jdxj",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestService_GetUser(t *testing.T) {
	rsp, err := client.GetUser(context.Background(), &user.GetUserReq{UserID: 1})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
