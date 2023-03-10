package main

import (
	"example/go-api/controllers"
	"example/go-api/initializers"
	"example/go-api/middlewares"
	"example/go-api/repositories"
	"example/go-api/services"
	"io"

	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	setupLogOutput()
}

var (
	movieRepostiory repositories.MovieRepostiory = repositories.NewMovieRepository()
	movieService    services.MovieService        = services.NewMovieService(movieRepostiory)
	movieController controllers.MoviesController = controllers.NewMovieController(movieService)

	reviewsRepostiory repositories.ReviewRepository = repositories.NewReviewRepository()
	reviewsService    services.ReviewService        = services.NewReviewService(reviewsRepostiory)
	reviewsController controllers.ReviewsController = controllers.NewReviewController(reviewsService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	r := gin.Default()
	defer CloseDB()

	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	movieGroup := r.Group("/Movies")
	{

		movieGroup.POST("/", movieController.CreateMovie)
		movieGroup.GET("/", movieController.GetMovies)
		movieGroup.GET("/:id", movieController.GetMovie)
		movieGroup.PUT("/:id", movieController.UpdateMovie)
		movieGroup.DELETE("/:id", movieController.DeleteMovie)
		movieGroup.GET("/:id/reviews", reviewsController.GetReviews)
		movieGroup.POST("/:id/reviews", reviewsController.AddReview)
	}
	reviewsGroup := r.Group("/Reviews")
	{
		reviewsGroup.GET("/:id", reviewsController.GetReviewById)
		reviewsGroup.PUT("/:id", reviewsController.UpdateReview)
		reviewsGroup.DELETE("/:id", reviewsController.DeleteReview)
	}
	r.Run()
}

func CloseDB() {
	dbConn, _ := initializers.DB.DB()
	dbConn.Close()
}
