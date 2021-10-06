package service

import (
	"context"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/crontab/dao/specification"
	"github.com/jdxj/sign/internal/crontab/dao/task"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func NewService() *Service {
	srv := &Service{
		cronParser: cron.NewParser(
			cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow),
	}
	return srv
}

type Service struct {
	crontab.UnimplementedCrontabServiceServer

	cronParser cron.Parser
}

func (srv *Service) CreateTask(ctx context.Context, req *crontab.CreateTaskReq) (*crontab.CreateTaskRsp, error) {
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}
	if req.Kind == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty kind")
	}
	_, err := srv.cronParser.Parse(req.Spec)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid cron spec")
	}
	if req.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid secret id")
	}

	return srv.createTask(req)
}

func (srv *Service) GetTasks(ctx context.Context, req *crontab.GetTasksReq) (*crontab.GetTasksRsp, error) {
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}

	where := map[string]interface{}{
		"user_id = ?": req.UserID,
	}
	if len(req.Kinds) != 0 {
		where["kind IN ?"] = req.Kinds
	}
	if len(req.SecretIDs) != 0 {
		where["secret_id IN ?"] = req.SecretIDs
	}

	tasks, err := task.Find(where)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get tasks failed: %s", err)
	}

	rsp := &crontab.GetTasksRsp{}
	for _, t := range tasks {
		pbTask := &crontab.Task{
			TaskID:   t.TaskID,
			UserID:   t.UserID,
			Describe: t.Describe,
			Kind:     crontab.Kind(t.Kind),
			Spec:     t.Spec,
			SecretID: t.SecretID,
		}
		rsp.List = append(rsp.List, pbTask)
	}
	return rsp, nil
}

func (srv *Service) DeleteTask(ctx context.Context, req *crontab.DeleteTaskReq) (*emptypb.Empty, error) {
	if req.TaskID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty task id")
	}
	err := task.Delete(map[string]interface{}{"task_id = ?": req.TaskID})
	return &emptypb.Empty{}, err
}

func (srv *Service) createTask(req *crontab.CreateTaskReq) (*crontab.CreateTaskRsp, error) {
	spec := &specification.Specification{
		Spec: req.Spec,
	}
	err := specification.Insert(spec)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert spec failed: %s", err)
	}

	t := &task.Task{
		Describe: req.Describe,
		UserID:   req.UserID,
		SecretID: req.SecretID,
		Kind:     int(req.Kind),
		Spec:     req.Spec,
	}
	err = task.Insert(t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert task failed: %s", err)
	}

	rsp := &crontab.CreateTaskRsp{
		TaskID: t.TaskID,
	}
	return rsp, nil
}
