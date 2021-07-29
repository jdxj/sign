package apiserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/jdxj/sign/internal/apiserver/v1"
	"github.com/jdxj/sign/internal/pkg/config"
)

func Run(conf config.APIServer) error {
	r := gin.Default()
	registerRouter(r, conf)

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	return r.Run(addr)
}

func registerRouter(r gin.IRouter, conf config.APIServer) {
	r.GET("/", hello)

	acc := gin.Accounts{conf.User: conf.Pass}
	api := r.Group("/api", gin.BasicAuth(acc))
	{
		v1.RegisterV1(api)
	}

}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}
