package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/app/handler"
)

func New(p gin.IRouter) {
	v1 := p.Group("/v1")
	v1.GET("/", handler.Hello)

	v1.POST("/login", handler.Login)
	v1.POST("/users", handler.SignUp)

	// user
	users := v1.Group("/users", authnBearer)
	{
		users.PUT("/:user_id", handler.UpdateUser)
	}

	// task
	tasks := v1.Group("/tasks", authnBearer)
	{
		tasks.POST("", handler.CreateTask)
		tasks.GET("", handler.GetTasks)
		tasks.GET("/:task_id", handler.GetTask)
		tasks.PUT("/:task_id", handler.UpdateTask)
		tasks.DELETE("/:task_id", handler.DeleteTask)
	}

}
