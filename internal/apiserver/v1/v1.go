package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver/common"
	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/code"
	"github.com/jdxj/sign/internal/task/bili"
	task "github.com/jdxj/sign/internal/task/common"
	"github.com/jdxj/sign/internal/task/hpi"
)

func RegisterV1(r gin.IRouter) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", hello)
		v1.POST("/task", addTask)
	}
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.NewResponse(0, "v1", nil))
}

func addTask(ctx *gin.Context) {
	req := &common.AddTaskReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.NewResponse(code.ErrBindReqFailed, err.Error(), nil))
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
		ctx.JSON(http.StatusBadRequest,
			common.NewResponse(code.ErrAuthFailed, err.Error(), nil))
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
			ctx.JSON(http.StatusBadRequest,
				common.NewResponse(code.ErrAddTaskFailed, err.Error(), nil))
			return
		}

		tmp := typ
		go func() {
			text := fmt.Sprintf("新任务被添加, id: %s, type: %s", req.ID, task.TypeMap[tmp])
			bot.Send(text)
		}()
	}
	ctx.JSON(http.StatusOK, common.NewResponse(0, "ok", nil))
}
