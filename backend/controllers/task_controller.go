package controllers

import (
	"net/http"
	"strconv"

	"github.com/carlkiptoo/backend/config"
	"github.com/carlkiptoo/backend/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	userID, _ := c.Get("user_id")
	config.DB.Where("user_id = ?", userID).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	var task models.Task
	userID, _ := c.Get("user_id")

	taskID := c.Param("id")

	result := config.DB.Where("user_id = ? AND id = ?", userID, taskID).First(&task)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	task.UserID = uint(userID.(int))


	config.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&task)
	c.JSON(http.StatusOK, task)

}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	config.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}