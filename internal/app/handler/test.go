package handler

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/app/api"
)

type HelloReq struct {
	Name string `json:"name" binding:"required"`
}

type HelloRsp struct {
	Name string `json:"name"`
}

func hello(_ context.Context, req *HelloReq) (*HelloRsp, error) {
	rsp := &HelloRsp{
		Name: fmt.Sprintf("hello %s", req.Name),
	}
	return rsp, nil
}

func Hello(ctx *gin.Context) {
	req := &HelloReq{}
	api.Process(ctx, req, func(request *api.Request) (interface{}, error) {
		return hello(ctx, req)
	})
}
