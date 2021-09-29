package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/pkg/code"
)

type rawRequest struct {
	Token   string      `json:"token"`
	Request interface{} `json:"request"`
}

type rawResponse struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Response    interface{} `json:"response"`
}

func Handle(ctx *gin.Context, req interface{}, f func(context.Context) (interface{}, error)) {
	rawReq := &rawRequest{
		Request: req,
	}
	err := ctx.Bind(rawReq)
	if err != nil {
		replyRawResponse(ctx, code.ErrBindReqFailed, err.Error(), nil)
		return
	}

	tCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rsp, err := f(tCtx)
	if err != nil {
		replyRawResponse(ctx, code.ErrHandle, err.Error(), nil)
		return
	}
	replyRawResponse(ctx, 0, "", rsp)
}

func replyRawResponse(ctx *gin.Context, code int, desc string, rsp interface{}) {
	rawRsp := &rawResponse{
		Code:        code,
		Description: desc,
		Response:    rsp,
	}
	ctx.JSON(http.StatusOK, rawRsp)
}
