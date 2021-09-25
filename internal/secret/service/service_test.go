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
	"github.com/jdxj/sign/internal/proto/secret"
)

var (
	client secret.SecretServiceClient
)

func TestMain(t *testing.M) {
	logger.Init("./secret.log")
	rpc.Init("127.0.0.1:2379")
	os.Exit(t.Run())
}

func TestService_CreateSecret(t *testing.T) {
	rpc.NewClient(secret.ServiceName, func(cc *grpc.ClientConn) {
		client = secret.NewSecretServiceClient(cc)
	})

	rsp, err := client.CreateSecret(context.Background(), &secret.CreateSecretReq{
		UserID: 1,
		Domain: crontab.Domain_BILI,
		Key:    key,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("secret id: %d\n", rsp.SecretID)

	rsp2, err := client.GetSecret(context.Background(), &secret.GetSecretReq{SecretID: rsp.SecretID})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp2)
}
