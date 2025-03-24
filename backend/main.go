package main

import (
	"github.com/carlkiptoo/backend/config"
	"github.com/carlkiptoo/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/carlkiptoo/backend/models"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Task{})

	routes.UserRoutes(r)
	routes.TaskRoutes(r)

	r.Run(":8080")
}