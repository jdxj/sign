package test_grpc

import (
	"context"
	"math/rand"

	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
)

type Service struct {
	testPB.UnimplementedTestRPCServer
}

func (s *Service) Hello(ctx context.Context, in *testPB.HelloReq) (*testPB.HelloRsp, error) {
	return &testPB.HelloRsp{
		Age: rand.Int63(),
	}, nil
}

type MultiService struct {
	testPB.UnimplementedTestMultiRPCServer
}

func (ms *MultiService) World(_ context.Context, req *testPB.WorldReq) (*testPB.WorldRsp, error) {
	return &testPB.WorldRsp{
		Age: rand.Int63(),
	}, nil
}
