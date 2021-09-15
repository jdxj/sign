package service

import (
	"context"

	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/crontab/dao/specification"
	"github.com/jdxj/sign/internal/crontab/dao/task"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func NewService() *Service {
	srv := &Service{
		cron: cron.New(),
		cronParser: cron.NewParser(
			cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow),
	}
	gPool, err := ants.NewPool(1000, ants.WithNonblocking(true))
	if err != nil {
		panic(err)
	}
	srv.gPool = gPool

	// todo: 从数据库中恢复 timer
	return srv
}

type Service struct {
	crontab.UnimplementedCrontabServiceServer

	cron       *cron.Cron
	cronParser cron.Parser

	gPool *ants.Pool
}

func (srv *Service) Start() {
	srv.cron.Start()
	logger.Infof("cron started")
}

func (srv *Service) Stop() {
	<-srv.cron.Stop().Done()
	srv.gPool.Release()
	logger.Infof("cron stopped")
}

func (srv *Service) CreateTask(ctx context.Context, req *crontab.CreateTaskRep) (*crontab.CreateTaskRsp, error) {
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}
	if req.Task == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty task define")
	}
	if _, ok := crontab.TaskKind_name[int32(req.Task.Kind)]; !ok {
		return nil, status.Errorf(codes.InvalidArgument, "kind not found: %d", req.Task.Kind)
	}
	_, err := srv.cronParser.Parse(req.Task.Spec)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid cron spec")
	}
	if req.Task.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid secret id")
	}

	return srv.createTask(req)
}

func (srv *Service) GetTasks(ctx context.Context, req *crontab.GetTasksReq) (*crontab.GetTasksRsp, error) {
	// todo: implement
	return nil, nil
}

func (srv *Service) UpdateTask(ctx context.Context, req *crontab.UpdateTaskReq) (*emptypb.Empty, error) {
	// todo: implement
	return nil, nil
}

func (srv *Service) DeleteTask(ctx context.Context, req *crontab.DeleteTaskReq) (*emptypb.Empty, error) {
	// todo: implement
	return nil, nil
}

func (srv *Service) createTask(req *crontab.CreateTaskRep) (*crontab.CreateTaskRsp, error) {
	spec := &specification.Specification{
		Spec: req.Task.Spec,
	}
	dup, err := specification.Insert(spec)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert spec failed: %s", err)
	}

	t := &task.Task{
		UserID:   req.UserID,
		Describe: req.Task.Describe,
		Kind:     int(req.Task.Kind),
		SpecID:   spec.SpecID,
		SecretID: req.Task.SecretID,
	}
	err = task.Insert(t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert task failed: %s", err)
	}

	// 只创建一个 spec 就行
	if dup == specification.NotDuplicateEntry {
		j := &job{
			specID: spec.SpecID,
			gPool:  srv.gPool,
		}
		_, err = srv.cron.AddJob(req.Task.Spec, j)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "add timer failed: %s", err)
		}
	} else {
		logger.Debugf("found duplicate spec, taskID: %d, specID: %d",
			t.TaskID, spec.SpecID)
	}

	rsp := &crontab.CreateTaskRsp{
		TaskID: t.TaskID,
	}
	return rsp, nil
}
