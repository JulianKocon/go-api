package main

import (
	"example/go-api/controllers"
	"example/go-api/initializers"
	"example/go-api/repositories"
	"example/go-api/services"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

var (
	movieRepostiory repositories.MovieRepostiory = repositories.NewMovieRepository()
	movieService    services.MovieService        = services.New(movieRepostiory)
	movieController controllers.MovieController  = controllers.New(movieService)
)

func main() {
	r := gin.Default()
	defer CloseDB()
	movieGroup := r.Group("/Movies")
	{

		movieGroup.POST("/CreateMovie", movieController.CreateMovie)
		movieGroup.GET("/GetMovies", movieController.GetMovies)
		movieGroup.GET("/GetMovie/:id", movieController.GetMovie)
		movieGroup.PUT("/UpdateMovie/:id", movieController.UpdateMovie)
		movieGroup.DELETE("/DeleteMovie/:id", movieController.DeleteMovie)
	}

	r.Run()
}

func CloseDB() {
	dbConn, _ := initializers.DB.DB()
	dbConn.Close()
}
