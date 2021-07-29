package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/jdxj/sign/pkg/apiserver/v1"
)

func register(r gin.IRouter) {
	r.GET("/", hello)
	r = r.Group("/api")
	v1.RegisterV1(r)
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}
