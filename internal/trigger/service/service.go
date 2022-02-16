package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/proto/trigger"
	"github.com/jdxj/sign/internal/trigger/dao"
)

func (s *Service) CreateTrigger(ctx context.Context, req *trigger.CreateTriggerRequest, _ *emptypb.Empty) error {
	spec := req.GetTrigger().GetSpec()
	_, err := s.parse(spec)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Parse: %s", err)
	}

	// 存在直接返回
	if s.hasAndAdd(spec) {
		logger.Debugf("spec: %s exists", spec)
		return nil
	}

	// 不存在则添加
	err = db.WithCtx(ctx).Create(&dao.Specification{
		Spec: spec,
	}).Error
	if err != nil {
		return status.Errorf(codes.Internal, "Create spec: %s", err)
	}

	err = s.addJob(spec, newJob(spec))
	if err != nil {
		return status.Errorf(codes.Internal, "AddJob: %s", err)
	}
	return nil
}

func (s *Service) GetTriggers(ctx context.Context, req *trigger.GetTriggersRequest, rsp *trigger.GetTriggersResponse) error {
	if req.GetOffset() < 0 || req.GetLimit() < 1 {
		return status.Errorf(codes.InvalidArgument, "invalid params")
	}

	err := db.WithCtx(ctx).
		Table(dao.TableName).
		Count(&rsp.Count).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	var rows []dao.Specification
	err = db.WithCtx(ctx).
		Order("spec_id").
		Offset(int(req.GetOffset())).
		Limit(int(req.GetLimit())).
		Find(&rows).Error
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, row := range rows {
		t := &trigger.Trigger{
			TriggerId: row.SpecID,
			Spec:      row.Spec,
		}
		rsp.Triggers = append(rsp.Triggers, t)
	}
	return nil
}
