package main

import (
	"github.com/carlkiptoo/backend/config"
	"github.com/carlkiptoo/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	routes.UserRoutes(r)
	routes.TaskRoutes(r)

	r.Run(":8080")
}