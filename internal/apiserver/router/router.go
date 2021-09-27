package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/jdxj/sign/internal/apiserver/v1"
)

func NewRouter() http.Handler {
	root := gin.Default()
	api := root.Group("/api")
	{
		v1.NewRouter(api)
	}
	return root
}
