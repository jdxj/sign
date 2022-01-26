package main

import (
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/pkg/util"
	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
	service "github.com/jdxj/sign/internal/test-grpc"
)

func main() {
	logger.Init("")

	server := rpc.NewServer(testPB.ServicePort)
	testPB.RegisterTestRPCServer(server, &service.Service{})
	server.Serve()

	util.Hold()

	server.Stop()
}
