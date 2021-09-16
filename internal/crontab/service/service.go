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
	"github.com/jdxj/sign/internal/crontab/model"
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
	t := req.Task
	if t == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty task define")
	}
	if t.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}
	if _, ok := crontab.TaskKind_name[int32(t.Kind)]; !ok {
		return nil, status.Errorf(codes.InvalidArgument, "kind not found: %d", req.Task.Kind)
	}
	_, err := srv.cronParser.Parse(t.Spec)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid cron spec")
	}
	if t.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid secret id")
	}

	return srv.createTask(t)
}

func (srv *Service) GetTasks(ctx context.Context, req *crontab.GetTasksReq) (*crontab.GetTasksRsp, error) {
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}

	return model.GetTasks(req)
}

func (srv *Service) DeleteTask(ctx context.Context, req *crontab.DeleteTaskReq) (*emptypb.Empty, error) {
	if req.TaskID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty task id")
	}
	err := task.Delete(map[string]interface{}{"task_id = ?": req.TaskID})
	return &emptypb.Empty{}, err
}

func (srv *Service) createTask(newTask *crontab.Task) (*crontab.CreateTaskRsp, error) {
	spec := &specification.Specification{
		Spec: newTask.Spec,
	}
	dup, err := specification.Insert(spec)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert spec failed: %s", err)
	}

	t := &task.Task{
		UserID:   newTask.UserID,
		Describe: newTask.Describe,
		Kind:     int(newTask.Kind),
		SpecID:   spec.SpecID,
		SecretID: newTask.SecretID,
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
		_, err = srv.cron.AddJob(newTask.Spec, j)
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
