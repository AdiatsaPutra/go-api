package controllers

import (
	"go-gin/configs"
	"go-gin/exception"
	"go-gin/models"
	"go-gin/payload"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := configs.DB.Create(&post)

	if result.Error != nil {
		exception.AppException(c, "Cant create post")
	}

	payload.HandleSuccess(c, post, "Post created", 200)
}

func GetPost(c *gin.Context) {
	var post []models.Post
	configs.DB.Find(&post)

	payload.HandleSuccess(c, post, "Post", 200)
}
