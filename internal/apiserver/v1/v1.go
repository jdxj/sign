package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver/comm"
	"github.com/jdxj/sign/internal/apiserver/model"
	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/code"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/task"
)

func RegisterV1(r gin.IRouter) {
	v1 := r.Group("/v1")
	v1.POST("/token", generateToken)

	test := v1.Group("/test", auth)
	{
		test.POST("/hello", testPost)
	}
	//{
	//	v1.GET("/", hello)
	//	v1.POST("/task", addTask)
	//
	//	v1.POST("/testPost", testPost)
	//}
}

func generateToken(ctx *gin.Context) {
	req := &comm.Request{}
	err := ctx.Bind(req)
	if err != nil {
		logger.Errorf("bind request-generate token err: %s", err)
		return
	}

	loginReq := &model.LoginReq{}
	err = json.Unmarshal(req.Data, loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			comm.NewResponse(
				code.ErrBindReqFailed,
				"unmarshal login request failed",
				err,
			),
		)
		return
	}

	resp, err := model.GenerateToken(loginReq)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			comm.NewResponse(
				code.ErrAuthFailed,
				"generate token failed",
				err,
			))
		return
	}
	ctx.JSON(http.StatusOK, comm.NewResponse(0, "ok", resp))
}

func auth(ctx *gin.Context) {
	req := &comm.Request{}
	err := ctx.Bind(req)
	if err != nil {
		logger.Errorf("auth, err: %s", err)
		return
	}

	claim, err := comm.CheckToken(req.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			comm.NewResponse(code.ErrAuthFailed, "check token failed", err),
		)
		return
	}
	ctx.Set("claim", claim)
	ctx.Set("data", string(req.Data))
}

func testPost(ctx *gin.Context) {
	req := &model.TestModule{}
	err := comm.UnmarshalRequest(ctx, req)
	if err != nil {
		logger.Errorf("testPost: %s", err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		comm.NewResponse(0, "hello: "+req.Nickname, nil),
	)
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, comm.NewResponse(0, "v1", nil))
}

func addTask(ctx *gin.Context) {
	req := &comm.AddTaskReq{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			comm.NewResponse(code.ErrBindReqFailed, err.Error(), nil))
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
			comm.NewResponse(code.ErrAddTaskFailed, err.Error(), nil))
		return
	}

	go func() {
		text := fmt.Sprintf("新任务被添加, id: %s, types: %v", t.ID, t.Types)
		bot.Send(text)
	}()
	ctx.JSON(http.StatusOK, comm.NewResponse(0, "ok", nil))
}
