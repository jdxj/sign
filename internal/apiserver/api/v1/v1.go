package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/apiserver"
	"github.com/jdxj/sign/internal/apiserver/model"
)

func NewRouter(parent gin.IRouter) {
	v1 := parent.Group("/v1", apiserver.Auth)
	// test
	{
		v1.GET("/hello", model.HandleHello)
	}
	// task
	{
		v1.POST("/task", model.CreateTask)
		v1.DELETE("/task", model.DeleteTask)
		v1.GET("/task", model.GetTasks)
	}
	// secret
	{
		v1.POST("/secret", model.CreateSecret)
		v1.PUT("/secret", model.UpdateSecret)
		v1.POST("/secret", model.GetSecret)
		v1.POST("/secrets", model.GetSecrets)
		v1.DELETE("/secret", model.DeleteSecret)
	}

}
