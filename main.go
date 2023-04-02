package main

import (
	// "go-gin/configs"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	configs.LoadEnvVariables()
// 	configs.ConnectDatabase()
// }

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
