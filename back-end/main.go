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
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Run()
}
