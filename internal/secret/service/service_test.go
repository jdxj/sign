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
	client secret.SecretServiceClient
)

func TestMain(t *testing.M) {
	logger.Init("./secret.log")
	rpc.Init("172.17.0.4:2379")
	rpc.NewClient(secret.ServiceName, func(cc *grpc.ClientConn) {
		client = secret.NewSecretServiceClient(cc)
	})

	os.Exit(t.Run())
}

func TestService_CreateSecret(t *testing.T) {
	rsp, err := client.CreateSecret(context.Background(), &secret.CreateSecretReq{
		UserID: 1,
		Domain: crontab.Domain_BILI,
		Key:    key,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("secret id: %d\n", rsp.SecretID)
}

func TestService_GetSecretList(t *testing.T) {
	rsp, err := client.GetSecretList(context.Background(), &secret.GetSecretListReq{
		UserID: 1,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range rsp.List {
		fmt.Printf("%+v\n", v)
	}
}
