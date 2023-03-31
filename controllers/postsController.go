package controllers

import (
	"github.com/brandonapol/go-crud/initializers"
	"github.com/brandonapol/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off request body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsGet(c *gin.Context) {
	// get data off request body

	var posts []models.Post

	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

