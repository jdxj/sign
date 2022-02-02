package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/jdxj/sign/internal/app/router/v1"
)

func New() http.Handler {
	root := gin.Default()
	api := root.Group("/api")
	{
		v1.New(api)
	}
	return root
}
