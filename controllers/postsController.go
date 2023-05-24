package controllers

import (
	"go-crud-gin/initializers"
	"go-crud-gin/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	ctx.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// Return post
	ctx.JSON(201, gin.H{
		"post": post,
	})
}

func GetAllPosts(ctx *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostById(ctx *gin.Context) {
	// Get id off url
	id := ctx.Param("id")
	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	// Get id off url
	id := ctx.Param("id")
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	ctx.Bind(&body)
	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)
	// Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(ctx *gin.Context) {
	// Get id off url
	id := ctx.Param("id")
	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)
	// Respond
	ctx.Status(204)
}
