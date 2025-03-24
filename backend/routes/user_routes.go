package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/carlkiptoo/backend/controllers"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
	}

}