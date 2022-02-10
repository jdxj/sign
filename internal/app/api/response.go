package api

import (
	"github.com/gin-gonic/gin"

	ser "github.com/jdxj/sign/internal/pkg/sign-error"
)

type Response struct {
	Code int         `json:"code,omitempty"`
	Desc string      `json:"desc,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Respond(ctx *gin.Context, data interface{}, err error) {
	var (
		rsp  = &Response{Data: data}
		code ser.Code
	)
	if err != nil {
		code = ser.ParseCode(err)
		rsp.Code = code.APP()
		rsp.Desc = code.String()
		rsp.Msg = err.Error()
	}
	ctx.JSON(code.HTTP(), rsp)
}
