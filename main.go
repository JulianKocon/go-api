package main

import (
	"example/go-api/controllers"
	"example/go-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	movieGroup := r.Group("/Movies")
	{
		movieGroup.POST("/CreateMovie", controllers.CreateMovie)
		movieGroup.GET("/GetMovies", controllers.GetMovies)
		movieGroup.GET("/GetMovie/:id", controllers.GetMovie)
		movieGroup.PUT("/UpdateMovie/:id", controllers.UpdateMovie)
		movieGroup.DELETE("/DeleteMovie/:id", controllers.DeleteMovie)
	}

	r.Run()
}
