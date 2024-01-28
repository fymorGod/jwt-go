package main

import (
	"github.com/fymorGod/jwt-go/controllers"
	"github.com/fymorGod/jwt-go/initializers"
	"github.com/fymorGod/jwt-go/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabases()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
