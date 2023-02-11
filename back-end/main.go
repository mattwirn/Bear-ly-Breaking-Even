package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/controllers"
	"github.com/mattwirn/Bear-ly-Breaking-Even/back-end/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabases()
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", controllers.Validate)
	router.Run()
}
