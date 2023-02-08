package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/models"
)

func Signup(c *gin.Context) {
	// Get the username/pass of req body

	// Hash the password

	// Create the user

	// Respond
	post := models.User{Username: "test", Password: "test"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": "post",
	})
}
