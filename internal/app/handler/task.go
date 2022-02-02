package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/jdxj/sign/internal/app/api"
	"github.com/jdxj/sign/internal/app/ref"
	ser "github.com/jdxj/sign/internal/pkg/sign-error"
	"github.com/jdxj/sign/internal/proto/task"
)

func paramToPB(kind string, body json.RawMessage) ([]byte, error) {
	msg := task.GetParamByKind(kind)
	if msg == nil {
		return nil, nil
	}
	if len(body) == 0 {
		return nil, nil
	}

	err := json.Unmarshal(body, msg)
	if err != nil {
		return nil, ser.New(ser.ErrInvalidParam, "参数转换失败: %s", err)
	}
	d, err := proto.Marshal(msg)
	if err != nil {
		return nil, ser.New(ser.ErrInternal, "参数转换失败: %s", err)
	}
	return d, nil
}

func pbToParam(kind string, body []byte) ([]byte, error) {
	msg := task.GetParamByKind(kind)
	if msg == nil {
		return nil, nil
	}
	if len(body) == 0 {
		return nil, nil
	}

	err := proto.Unmarshal(body, msg)
	if err != nil {
		return nil, ser.New(ser.ErrInternal, "Unmarshal: %s", err)
	}
	body, err = json.Marshal(msg)
	if err != nil {
		return nil, ser.New(ser.ErrInternal, "Marshal: %s", err)
	}
	return body, nil
}

type CreateTaskReq struct {
	Desc string `json:"desc"`
	Kind string `json:"kind" binding:"required"`
	Spec string `json:"spec" binding:"required"`

	// Param 不同的 kind 有不同的 param, 所以这里不对 Param 解析
	Param json.RawMessage `json:"param"`
}

type CreateTaskRsp struct {
	TaskID int64 `json:"task_id"`
}

func createTask(ctx context.Context, req *CreateTaskReq, userID int64) (*CreateTaskRsp, error) {
	param, err := paramToPB(req.Kind, req.Param)
	if err != nil {
		return nil, err
	}

	ctRsp, err := ref.TaskService.CreateTask(ctx, &task.CreateTaskRequest{Task: &task.Task{
		Description: req.Desc,
		UserId:      userID,
		Kind:        req.Kind,
		Spec:        req.Spec,
		Param:       param,
	}})
	if err != nil {
		return nil, ser.New(ser.ErrRPCRequest, "CreateTask: %s", err)
	}

	rsp := &CreateTaskRsp{
		TaskID: ctRsp.GetTaskId(),
	}
	return rsp, nil
}

func CreateTask(ctx *gin.Context) {
	req := &CreateTaskReq{}
	api.ProcessCheckToken(ctx, req, func(request *api.Request) (interface{}, error) {
		return createTask(ctx, req, request.Claim.UserID)
	})
}

type GetTaskReq struct {
	TaskID int64 `json:"task_id" binding:"required"`
}

type Task struct {
	TaskID    int64  `json:"task_id"`
	Desc      string `json:"desc"`
	UserID    int64  `json:"user_id"`
	Kind      string `json:"kind"`
	Spec      string `json:"spec"`
	CreatedAt int64  `json:"created_at"`

	// Param 返回时也不要进行 base64 编码
	Param json.RawMessage `json:"param"`
}

type GetTaskRsp struct {
	Task Task `json:"task"`
}

func getTask(ctx context.Context, req *GetTaskReq, userID int64) (*GetTaskRsp, error) {
	gtRsp, err := ref.TaskService.GetTask(ctx, &task.GetTaskRequest{
		TaskId: req.TaskID,
		UserId: userID,
	})
	if err != nil {
		return nil, ser.New(ser.ErrRPCRequest, "GetTask: %s", err)
	}

	t := gtRsp.GetTask()
	param, err := pbToParam(t.GetKind(), t.GetParam())
	if err != nil {
		return nil, err
	}

	rsp := &GetTaskRsp{
		Task: Task{
			TaskID:    t.GetTaskId(),
			Desc:      t.GetDescription(),
			UserID:    t.GetUserId(),
			Kind:      t.GetKind(),
			Spec:      t.GetSpec(),
			Param:     param,
			CreatedAt: t.GetCreatedAt().AsTime().Unix(),
		}}
	return rsp, nil
}

func GetTask(ctx *gin.Context) {
	req := &GetTaskReq{}
	api.ProcessCheckToken(ctx, req, func(request *api.Request) (interface{}, error) {
		return getTask(ctx, req, request.Claim.UserID)
	})
}

type GetTasksReq struct {
	TaskID    int64  `json:"task_id"`
	Desc      string `json:"desc"`
	Kind      string `json:"kind"`
	Spec      string `json:"spec"`
	CreatedAt int64  `json:"created_at,string"`
	PageID    int64  `json:"page_id,string" binding:"gt=0"`
	PageSize  int64  `json:"page_size,string" binding:"gt=0"`
}

type GetTasksRsp struct {
	Count int64  `json:"count"`
	Tasks []Task `json:"tasks"`
}

func getTasks(ctx context.Context, req *GetTasksReq, userID int64) (*GetTasksRsp, error) {
	gtRsp, err := ref.TaskService.GetTasks(ctx, &task.GetTasksRequest{
		TaskId:      req.TaskID,
		Description: req.Desc,
		UserId:      userID,
		Kind:        req.Kind,
		Spec:        req.Spec,
		CreatedAt:   timestamppb.New(time.Unix(req.CreatedAt, 0)),
		Offset:      (req.PageID - 1) * req.PageSize,
		Limit:       req.PageSize,
	})
	if err != nil {
		return nil, ser.New(ser.ErrRPCRequest, "GetTasks: %s", err)
	}

	rsp := &GetTasksRsp{
		Count: gtRsp.GetCount(),
	}
	for _, v := range gtRsp.GetTasks() {
		param, err := pbToParam(v.GetKind(), v.GetParam())
		if err != nil {
			return nil, err
		}
		t := Task{
			TaskID:    v.GetTaskId(),
			Desc:      v.GetDescription(),
			UserID:    v.GetUserId(),
			Kind:      v.GetKind(),
			Spec:      v.GetSpec(),
			Param:     param,
			CreatedAt: v.GetCreatedAt().AsTime().Unix(),
		}
		rsp.Tasks = append(rsp.Tasks, t)
	}
	return rsp, nil
}

func GetTasks(ctx *gin.Context) {
	req := &GetTasksReq{}
	api.ProcessCheckToken(ctx, req, func(request *api.Request) (interface{}, error) {
		return getTasks(ctx, req, request.Claim.UserID)
	})
}

type UpdateTaskReq struct {
	TaskID int64  `json:"task_id,string" binding:"required"`
	Desc   string `json:"desc"`
	Spec   string `json:"spec"`

	// Param 更新时也不要解析
	Param json.RawMessage `json:"param"`
}

func updateTask(ctx context.Context, req *UpdateTaskReq, userID int64) error {
	gtRsp, err := ref.TaskService.GetTask(ctx, &task.GetTaskRequest{
		TaskId: req.TaskID,
		UserId: userID,
	})
	if err != nil {
		return ser.New(ser.ErrRPCRequest, "GetTask: %s", err)
	}
	param, err := paramToPB(gtRsp.GetTask().GetKind(), req.Param)
	if err != nil {
		return err
	}

	_, err = ref.TaskService.UpdateTask(ctx, &task.UpdateTaskRequest{Task: &task.Task{
		TaskId:      req.TaskID,
		Description: req.Desc,
		UserId:      userID,
		Spec:        req.Spec,
		Param:       param,
	}})
	if err != nil {
		return ser.New(ser.ErrRPCRequest, "UpdateTask: %s", err)
	}
	return nil
}

func UpdateTask(ctx *gin.Context) {
	req := &UpdateTaskReq{}
	api.ProcessCheckToken(ctx, req, func(request *api.Request) (interface{}, error) {
		return nil, updateTask(ctx, req, request.Claim.UserID)
	})
}

type DeleteTaskReq struct {
	TaskID int64 `json:"task_id" binding:"required"`
}

func deleteTask(ctx context.Context, req *DeleteTaskReq, userID int64) error {
	_, err := ref.TaskService.DeleteTask(ctx, &task.DeleteTaskRequest{
		TaskId: req.TaskID,
		UserId: userID,
	})
	if err != nil {
		return ser.New(ser.ErrRPCRequest, "DeleteTask: %s", err)
	}
	return nil
}

func DeleteTask(ctx *gin.Context) {
	req := &DeleteTaskReq{}
	api.ProcessCheckToken(ctx, req, func(request *api.Request) (interface{}, error) {
		return nil, deleteTask(ctx, req, request.Claim.UserID)
	})
}
