package main

import (
	"go-crud-gin/controllers"
	"go-crud-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVarialbles()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run()
}
