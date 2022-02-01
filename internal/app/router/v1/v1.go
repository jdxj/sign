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

	// session
	session := v1.Group("/session")
	{
		session.POST("/login", handler.Login)
	}

	// user
	user := v1.Group("/user")
	{
		user.POST("/sign-up", handler.SignUp)
		user.POST("/update", handler.UpdateUser)
	}

	// task
	task := v1.Group("/task")
	{
		task.POST("/create", handler.CreateTask)
		task.POST("/get", handler.GetTask)
		task.POST("list", handler.GetTasks)
		task.POST("/update", handler.UpdateTask)
		task.POST("/delete", handler.DeleteTask)
	}
}
