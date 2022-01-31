package main

import (
	"net/rpc"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/util"
	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
	service "github.com/jdxj/sign/internal/test-grpc"
)

func main() {
	logger.Init("")

	multiS := rpc.NewServer(testPB.MServicePort)
	testPB.RegisterTestMultiRPCServer(multiS, &service.MultiService{})
	multiS.Serve()

	util.Hold()

	multiS.Stop()
}
