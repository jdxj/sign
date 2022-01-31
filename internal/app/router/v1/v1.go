package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/app/handler"
)

func New(p gin.IRouter) {
	v1 := p.Group("/v1")
	// test
	{
		v1.POST("/hello", handler.Hello)
	}
	// task
	{
		//v1.POST("/task", model.CreateTask)
		//v1.DELETE("/task", model.DeleteTask)
		//v1.POST("/tasks", model.GetTasks)
	}
}
