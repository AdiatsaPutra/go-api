package main

import (
	"go-gin/configs"
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/post", controllers.GetPost)
	r.POST("/post", controllers.PostCreate)
	r.Run()
}
