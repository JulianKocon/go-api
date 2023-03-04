package main

import (
	"example/go-api/initializers"
	"example/go-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Movie{})
}
