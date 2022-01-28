package main

import (
	"go-micro.dev/v4"

	"github.com/jdxj/sign/internal/pkg/logger"
	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
	impl "github.com/jdxj/sign/internal/test-grpc"
)

func main() {
	logger.Init("")

	service := micro.NewService(
		micro.Name("test-grpc"),
	)
	service.Init()
	_ = testPB.RegisterTestRPCHandler(service.Server(), new(impl.Service))

	err := service.Run()
	if err != nil {
		logger.Errorf("run: %s", err)
	}
}
