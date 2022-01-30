package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/proto/trigger"
	"github.com/jdxj/sign/internal/trigger/dao"
)

func (s *Service) CreateTrigger(ctx context.Context, req *trigger.CreateTriggerRequest, _ *emptypb.Empty) error {
	spec := req.GetTrigger().GetSpec()
	_, err := s.parser.Parse(spec)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Parse: %s", err)
	}

	// 存在直接返回
	if s.hasAndAdd(spec) {
		return nil
	}

	// 不存在则添加
	err = db.WithCtx(ctx).Create(&dao.Specification{
		Spec: spec,
	}).Error
	if err != nil {
		return status.Errorf(codes.Internal, "Create spec: %s", err)
	}

	_, err = s.cron.AddJob(spec, newJob(spec))
	if err != nil {
		return status.Errorf(codes.Internal, "AddJob: %s", err)
	}
	return nil
}

func (s *Service) GetTriggers(ctx context.Context, req *trigger.GetTriggersRequest, rsp *trigger.GetTriggersResponse) error {
	// todo: implement
	return status.Errorf(codes.Unimplemented, "GetTriggers")
}
