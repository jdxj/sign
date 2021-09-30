package v1

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/apiserver/model"
)

func NewRouter(parent gin.IRouter) {
	v1 := parent.Group("/v1", apiserver.Auth)
	v1.POST("/hello", testPost)
}

func testPost(ctx *gin.Context) {
	req := &model.Hello{}
	apiserver.Handle(ctx, req, func(tCtx context.Context) (interface{}, error) {
		return model.HandleHello(tCtx, req)
	})
}
