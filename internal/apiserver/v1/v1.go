package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver/common"
	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/code"
	"github.com/jdxj/sign/internal/task"
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

	t := &task.Task{
		ID:     req.ID,
		Domain: req.Domain,
		Types:  req.Type,
		Key:    req.Key,
	}
	err = task.Add(t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.NewResponse(code.ErrAddTaskFailed, err.Error(), nil))
		return
	}

	go func() {
		text := fmt.Sprintf("新任务被添加, id: %s, types: %v", t.ID, t.Types)
		bot.Send(text)
	}()
	ctx.JSON(http.StatusOK, common.NewResponse(0, "ok", nil))
}
