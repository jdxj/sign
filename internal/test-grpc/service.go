package test_grpc

import (
	"context"
	"math/rand"

	testPB "github.com/jdxj/sign/internal/proto/test-grpc"
)

type Service struct {
}

func (s *Service) Hello(ctx context.Context, req *testPB.HelloReq, rsp *testPB.HelloRsp) error {
	rsp.Age = rand.Int63()
	return nil
}
