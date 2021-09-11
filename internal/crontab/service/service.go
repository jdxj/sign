package service

import (
	"context"

	"github.com/jdxj/sign/internal/proto/crontab"
)

type Service struct {
	crontab.UnimplementedTestServiceServer
}

func (srv *Service) Hello(ctx context.Context, req *crontab.TestReq) (*crontab.TestRsp, error) {
	rsp := &crontab.TestRsp{
		Reply: "hello " + req.Nickname,
	}
	return rsp, nil
}
