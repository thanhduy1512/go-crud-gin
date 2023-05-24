package main

import (
	"go-crud-gin/initializers"
	"go-crud-gin/models"
)

func init() {
	initializers.LoadEnvVarialbles()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
