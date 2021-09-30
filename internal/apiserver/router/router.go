package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver"
	v1 "github.com/jdxj/sign/internal/apiserver/api/v1"
)

func New() http.Handler {
	root := gin.Default()
	api := root.Group("/api")
	{
		api.POST("/token", apiserver.Login)
		v1.NewRouter(api)
	}
	return root
}
