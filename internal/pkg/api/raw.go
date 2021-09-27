package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RawRequest struct {
	Token   string      `json:"token"`
	Request interface{} `json:"request"`
}

type RawResponse struct {
	ErrorCode        int         `json:"error_code"`
	ErrorDescription string      `json:"error_description"`
	Response         interface{} `json:"response"`
}

func ParseRawRequest(ctx *gin.Context, req interface{}) (*RawRequest, error) {
	rawReq := &RawRequest{
		Request: req,
	}
	return rawReq, ctx.Bind(rawReq)
}

func ReplyRawResponse(ctx *gin.Context, rsp interface{}, err error) {
	rawRsp := &RawResponse{
		ErrorCode:        0,
		ErrorDescription: "",
		Response:         rsp,
	}
	if err != nil {
		// todo: 实现自定义 Error
		rawRsp.ErrorCode = 1
		rawRsp.ErrorDescription = err.Error()
	}
	ctx.JSON(http.StatusOK, rawRsp)
}
