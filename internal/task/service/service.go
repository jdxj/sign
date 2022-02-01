package service

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/proto/trigger"
	"github.com/jdxj/sign/internal/task/dao"
)

func New(conf config.Secret) *Service {
	srv := &Service{
		key: []byte(conf.Key),
	}
	return srv
}

type Service struct {
	key []byte
}

func (s *Service) CreateTask(ctx context.Context, req *pb.CreateTaskRequest, rsp *pb.CreateTaskResponse) error {
	task := req.GetTask()
	if task.GetUserId() == 0 {
		return status.Errorf(codes.InvalidArgument, "userId is empty")
	}

	if _, ok := pb.Kind_value[task.Kind]; !ok || task.GetKind() == pb.Kind_UNKNOWN_KIND.String() {
		return status.Errorf(codes.Internal, "invalid kind")
	}

	_, err := triggerService.CreateTrigger(ctx, &trigger.CreateTriggerRequest{Trigger: &trigger.Trigger{
		Spec: task.GetSpec(),
	}})
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "CreateTrigger: %s", err)
	}

	row := dao.Task{
		Description: task.GetDescription(),
		UserID:      task.GetUserId(),
		Kind:        task.GetKind(),
		Spec:        task.GetSpec(),
		Param:       util.Encrypt(s.key, task.GetParam()),
	}
	err = db.WithCtx(ctx).
		Select("description", "user_id", "kind", "spec", "param").
		Create(&row).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, "Create: %s", err)
	}
	rsp.TaskId = row.TaskID
	return nil
}

func (s *Service) GetTask(ctx context.Context, req *pb.GetTaskRequest, rsp *pb.GetTaskResponse) error {
	if req.GetTaskId() == 0 || req.GetUserId() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid param")
	}

	var row dao.Task
	err := db.WithCtx(ctx).
		Where("task_id = ?", req.GetTaskId()).
		Where("user_id = ?", req.GetUserId()).
		Take(&row).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, "Take: %s", err)
	}

	rsp.Task = &pb.Task{
		TaskId:      row.TaskID,
		Description: row.Description,
		UserId:      row.UserID,
		Kind:        row.Kind,
		Spec:        row.Spec,
		Param:       util.Decrypt(s.key, row.Param),
		CreatedAt:   timestamppb.New(row.CreatedAt),
	}
	return nil
}

func (s *Service) GetTasks(ctx context.Context, req *pb.GetTasksRequest, rsp *pb.GetTasksResponse) error {
	if req.GetOffset() < 0 || req.GetLimit() < 1 {
		return status.Errorf(codes.InvalidArgument, "invalid pagination")
	}

	db := db.WithCtx(ctx).
		Model(&dao.Task{})
	if req.GetTaskId() != 0 {
		db.Where("task_id = ?", req.GetTaskId())
	}
	if req.GetDescription() != "" {
		db.Where("description LIKE ?", fmt.Sprintf("%%%s%%", req.GetDescription()))
	}
	if req.GetUserId() != 0 {
		db.Where("user_id = ?", req.GetUserId())
	}
	if req.GetKind() != "" {
		db.Where("kind = ?", req.GetKind())
	}
	if req.GetSpec() != "" {
		db.Where("spec = ?", req.GetSpec())
	}
	if req.GetCreatedAt().IsValid() {
		db.Where("created_at >= ?", req.GetCreatedAt().AsTime())
	}

	err := db.Count(&rsp.Count).Error
	if err != nil {
		return status.Errorf(codes.Internal, "Count: %s", err)
	}

	var rows []dao.Task
	err = db.Order("created_at DESC, task_id DESC").
		Offset(int(req.GetOffset())).
		Limit(int(req.GetLimit())).
		Find(&rows).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, "Find: %s", err)
	}

	for _, row := range rows {
		t := &pb.Task{
			TaskId:      row.TaskID,
			Description: row.Description,
			UserId:      row.UserID,
			Kind:        row.Kind,
			Spec:        row.Spec,
			Param:       util.Decrypt(s.key, row.Param),
			CreatedAt:   timestamppb.New(row.CreatedAt),
		}
		rsp.Tasks = append(rsp.Tasks, t)
	}
	return nil
}

func (s *Service) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest, _ *emptypb.Empty) error {
	task := req.GetTask()
	if task.GetTaskId() == 0 {
		return status.Errorf(codes.InvalidArgument, "taskId is empty")
	}

	data := make(map[string]interface{})
	if task.GetDescription() != "" {
		data["description"] = task.GetDescription()
	}
	if task.GetSpec() != "" {
		_, err := triggerService.CreateTrigger(ctx, &trigger.CreateTriggerRequest{Trigger: &trigger.Trigger{
			Spec: task.GetSpec(),
		}})
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "CreateTrigger: %s", err)
		}
		data["spec"] = task.GetSpec()
	}
	if len(task.GetParam()) != 0 {
		data["param"] = util.Encrypt(s.key, task.GetParam())
	}
	if len(data) == 0 {
		return nil
	}

	err := db.WithCtx(ctx).
		Model(&dao.Task{}).
		Where("task_id = ?", task.GetTaskId()).
		Updates(data).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, "Updates: %s", err)
	}
	return nil
}

func (s *Service) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest, _ *emptypb.Empty) error {
	if req.GetTaskId() == 0 || req.GetUserId() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid param")
	}

	err := db.WithCtx(ctx).
		Where("task_id = ?", req.GetTaskId()).
		Where("user_id = ?", req.GetUserId()).
		Delete(&dao.Task{}).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, "Delete: %s", err)
	}
	return nil
}

func (s *Service) DispatchTasks(_ context.Context, req *pb.DispatchTasksRequest, _ *emptypb.Empty) error {
	err := gPool.Submit(newJob(s.key, req.GetSpec()).dispatchTasks)
	if err != nil {
		return status.Errorf(codes.Internal, "Submit: %s", err)
	}
	return nil
}
