package main

import (
	"go-gin/configs"
	"go-gin/models"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectDatabase()
}

func main() {
	configs.DB.AutoMigrate(&models.Post{})
}
