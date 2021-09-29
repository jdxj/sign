package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/apiserver/comm"
	"github.com/jdxj/sign/internal/apiserver/model"
	"github.com/jdxj/sign/internal/pkg/api"
	"github.com/jdxj/sign/internal/pkg/code"
	"github.com/jdxj/sign/internal/pkg/logger"
)

func NewRouter(parent gin.IRouter) {
	v1 := parent.Group("/v1", apiserver.Auth)
	v1.POST("/hello", testPost)
}

func generateToken(ctx *gin.Context) {
	req := &model.LoginReq{}
	_, err := api.ParseRawRequest(ctx, req)
	if err != nil {
		api.ReplyRawResponse(ctx, code.ErrBindReqFailed, err.Error(), nil)
		return
	}

	tCtx, cancel := comm.TimeoutContext()
	defer cancel()

	rsp, err := model.GenerateToken(tCtx, req)
	if err != nil {
		api.ReplyRawResponse(ctx, code.ErrAuthFailed, err.Error(), nil)
		return
	}
	api.ReplyRawResponse(ctx, 0, "", rsp)
}

func auth(ctx *gin.Context, req interface{}) error {
	rawReq, err := api.ParseRawRequest(ctx, req)
	if err != nil {
		return err
	}
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
