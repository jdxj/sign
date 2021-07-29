package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	task "github.com/jdxj/sign/pkg/task/common"
	"github.com/jdxj/sign/pkg/task/hpi"

	"github.com/jdxj/sign/pkg/apiserver/common"
	"github.com/jdxj/sign/pkg/task/bili"
)

func RegisterV1(r gin.IRouter) {
	r = r.Group("/v1")
	r.POST("/task", addTask)
}

func addTask(ctx *gin.Context) {
	req := &common.AddTaskReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &common.AddTaskResp{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	var (
		client *http.Client
	)
	switch req.Domain {
	case task.BiliDomain:
		client, err = bili.Auth(req.Key)

	case task.HPIDomain:
		client, err = hpi.Auth(req.Key)

	default:
		err = fmt.Errorf("unsupport domain: %d", req.Domain)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, &common.AddTaskResp{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	for _, typ := range req.Type {
		t := &task.Task{
			ID:     req.ID,
			Type:   typ,
			Client: client,
		}

		switch typ {
		case task.BiliSign:
			bili.AddSignTask(t)

		case task.BiliBCount:
			bili.AddBCountTask(t)

		case task.HPISign:
			hpi.AddSignTask(t)

		default:
			err = fmt.Errorf("unsupport type: %d", typ)
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, &common.AddTaskResp{
				Code:    1,
				Message: err.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, &common.AddTaskResp{
		Code:    0,
		Message: "",
	})
}
