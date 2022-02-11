package v1

import (
	"github.com/gin-gonic/gin"
)

func New(p gin.IRouter) {
	v1 := p.Group("/v1")

	v1.POST("/login", Login)
	v1.POST("/users", SignUp)

	// user
	users := v1.Group("/users", authnBearer)
	{
		users.PUT("/:user_id", UpdateUser)
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
