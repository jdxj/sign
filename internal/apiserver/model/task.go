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

func HandleHello(ctx context.Context, req *Hello) (*World, error) {
	rsp := &World{
		Reply: fmt.Sprintf("hello %s!", req.Nickname),
	}
	return rsp, nil
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
