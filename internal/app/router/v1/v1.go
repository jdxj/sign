package v1

import (
	"github.com/gin-gonic/gin"
)

func New(p gin.IRouter) {
	v1 := p.Group("/v1")

	// user 无认证
	noAuthn := v1.Group("/users")
	{
		noAuthn.POST("/user", SignUp)
		noAuthn.POST("/token", Login)
	}

	// user 有认证
	users := v1.Group("/users", authnBearer)
	{
		users.PUT("/user", UpdateUser)
	}

	// task
	tasks := v1.Group("/tasks", authnBearer)
	{
		tasks.POST("", CreateTask)
		tasks.GET("", GetTasks)
		tasks.GET("/:task_id", GetTask)
		tasks.PUT("/:task_id", UpdateTask)
		tasks.DELETE("/:task_id", DeleteTask)
	}
}
