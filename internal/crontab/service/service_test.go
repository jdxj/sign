package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/proto/crontab"
)

func TestService_Hello(t *testing.T) {
	cc, err := grpc.Dial("127.0.0.1:49152", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer cc.Close()

	client := crontab.NewTestServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &crontab.TestReq{
		Nickname: "jdxj",
	}
	rsp, err := client.Hello(ctx, req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}
