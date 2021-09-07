package common

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/pkg/code"
)

type Request struct {
	Token string          `json:"token"`
	Data  json.RawMessage `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AddTaskReq struct {
	ID     string `json:"id"`
	Domain int    `json:"domain"`
	Type   []int  `json:"type"`
	Key    string `json:"key"`
}

func UnmarshalRequest(ctx *gin.Context, req interface{}) error {
	data := ctx.GetString("data")
	err := json.Unmarshal([]byte(data), req)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			NewResponse(
				code.ErrBindReqFailed,
				"unmarshal request failed",
				err))
	}
	return err
}
