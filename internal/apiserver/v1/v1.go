package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver/common"
	"github.com/jdxj/sign/internal/apiserver/module"
	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/code"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/task"
)

func RegisterV1(r gin.IRouter) {
	v1 := r.Group("/v1", auth)
	{
		v1.GET("/", hello)
		v1.POST("/task", addTask)

		v1.POST("/testPost", testPost)
	}
}

func auth(ctx *gin.Context) {
	req := &common.Request{}
	err := ctx.Bind(req)
	if err != nil {
		logger.Errorf("auth, err: %s", err)
		return
	}

	// todo: 使用 redis
	if req.Token != "jdxj" {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			common.NewResponse(code.ErrAuthFailed, "invalid token", nil),
		)
		return
	}
	ctx.Set("data", string(req.Data))
}

func testPost(ctx *gin.Context) {
	req := &module.TestModule{}
	err := common.UnmarshalRequest(ctx, req)
	if err != nil {
		logger.Errorf("testPost: %s", err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.NewResponse(
			code.Hello,
			"hello: "+req.Nickname,
			nil,
		))
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
