package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver/comm"
	"github.com/jdxj/sign/internal/apiserver/model"
	"github.com/jdxj/sign/internal/pkg/code"
	"github.com/jdxj/sign/internal/pkg/logger"
)

func NewRouter(parent gin.IRouter) {
	v1 := parent.Group("/v1")
	v1.POST("/token", generateToken)

	test := v1.Group("/test", auth)
	{
		test.POST("/hello", testPost)
	}
}

func generateToken(ctx *gin.Context) {
	req := &comm.Request{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			comm.NewResponse(code.ErrBindReqFailed, "bind comm request failed", nil),
		)
		logger.Errorf("bind comm request failed: %s", err)
		return
	}

	loginReq := &model.LoginReq{}
	err = json.Unmarshal(req.Data, loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			comm.NewResponse(code.ErrBindReqFailed, "unmarshal login request failed", nil),
		)
		logger.Errorf("unmarshal login req failed: %s", err)
		return
	}

	tCtx, cancel := comm.TimeoutContext()
	defer cancel()

	resp, err := model.GenerateToken(tCtx, loginReq)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			comm.NewResponse(code.ErrAuthFailed, "generate token failed", err.Error()),
		)
		logger.Errorf("generate token failed: %s", err)
		return
	}
	ctx.JSON(http.StatusOK, comm.NewResponse(0, "ok", resp))
}

func auth(ctx *gin.Context) {
	req := &comm.Request{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			comm.NewResponse(code.ErrBindReqFailed, "bind comm request failed", nil),
		)
		logger.Errorf("bind comm request failed: %s", err)
		return
	}

	claim, err := comm.CheckToken(req.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			comm.NewResponse(code.ErrAuthFailed, "check token failed", nil),
		)
		logger.Errorf("check token failed: %s", err)
		return
	}
	ctx.Set("claim", claim)
	ctx.Set("data", string(req.Data))
}

func testPost(ctx *gin.Context) {
	req := &model.LoginReq{}
	err := comm.UnmarshalRequest(ctx, req)
	if err != nil {
		logger.Errorf("unmarshal request failed: %s", err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		comm.NewResponse(0, "hello: "+req.Nickname, nil),
	)
}
