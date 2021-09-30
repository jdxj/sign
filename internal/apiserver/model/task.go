package model

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/proto/crontab"
)

type Hello struct {
	Nickname string `json:"nickname"`
}

type World struct {
	Reply string `json:"reply"`
}

func HandleHello(ctx *gin.Context) {
	req := &Hello{}
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return func() (interface{}, error) {
			rsp := &World{
				Reply: fmt.Sprintf("hello %s!", req.Nickname),
			}
			return rsp, nil
		}()
	})
}

type CreateTaskReq struct {
	UserID   int64        `json:"user_id"`
	Describe string       `json:"describe"`
	Kind     crontab.Kind `json:"kind"`
	Spec     string       `json:"spec"`
	SecretID int64        `json:"secret_id"`
}
type CreateTaskRsp struct {
	TaskID int64 `json:"task_id"`
}

func CreateTask(ctx *gin.Context) {
	req := &CreateTaskReq{}
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return createTask(tCtx, req)
	})
}

func createTask(ctx context.Context, req *CreateTaskReq) (*CreateTaskRsp, error) {
	createRsp, err := apiserver.CronClient.CreateTask(ctx, &crontab.CreateTaskReq{
		UserID:   req.UserID,
		Describe: req.Describe,
		Kind:     req.Kind,
		Spec:     req.Spec,
		SecretID: req.SecretID,
	})
	if err != nil {
		return nil, err
	}

	rsp := &CreateTaskRsp{
		TaskID: createRsp.TaskID,
	}
	return rsp, nil
}

type DeleteTaskReq struct {
	TaskID int64 `json:"task_id"`
}

func DeleteTask(ctx *gin.Context) {
	req := &DeleteTaskReq{}
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return nil, deleteTask(tCtx, req)
	})
}

func deleteTask(ctx context.Context, req *DeleteTaskReq) error {
	_, err := apiserver.CronClient.DeleteTask(ctx, &crontab.DeleteTaskReq{
		TaskID: req.TaskID,
	})
	return err
}

type GetTasksReq struct {
	UserID int64 `json:"user_id"`
}

type Task struct {
	TaskID   int64        `json:"task_id"`
	Describe string       `json:"describe"`
	Kind     crontab.Kind `json:"kind"`
	Spec     string       `json:"spec"`
	SecretID int64        `json:"secret_id"`
}

type GetTasksRsp struct {
	List []*Task `json:"list"`
}

func GetTasks(ctx *gin.Context) {
	req := &GetTasksReq{}
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return getTasks(tCtx, req)
	})
}

func getTasks(ctx context.Context, req *GetTasksReq) (*GetTasksRsp, error) {
	tasks, err := apiserver.CronClient.GetTasks(ctx, &crontab.GetTasksReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	rsp := &GetTasksRsp{}
	for _, v := range tasks.List {
		t := &Task{
			TaskID:   v.TaskID,
			Describe: v.Describe,
			Kind:     v.Kind,
			Spec:     v.Spec,
			SecretID: v.SecretID,
		}
		rsp.List = append(rsp.List, t)
	}
	return rsp, nil
}
