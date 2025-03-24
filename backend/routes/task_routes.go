package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/carlkiptoo/backend/controllers"
	"github.com/carlkiptoo/backend/middlewares"
)

func TaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/tasks")
	taskGroup.Use(middlewares.AuthMiddleware())
	{
		taskGroup.GET("/", controllers.GetTasks)
		taskGroup.GET("/:id", controllers.GetTaskById)
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}
}