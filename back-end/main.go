package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
