package controllers

import (
	"github.com/carlkiptoo/backend/config"
	"net/http"
	"os"
	"time"

	"github.com/carlkiptoo/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}