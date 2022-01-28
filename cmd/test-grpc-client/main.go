package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v4"

	test_grpc "github.com/jdxj/sign/internal/proto/test-grpc"
)

func main() {
	service := micro.NewService()
	service.Init()
	client := test_grpc.NewTestRPCService("test-grpc", service.Client())
	rsp, err := client.Hello(context.Background(), &test_grpc.HelloReq{
		Name: "abc",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", rsp)
}
