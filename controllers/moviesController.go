package controllers

import (
	"example/go-api/initializers"
	"example/go-api/models"

	"github.com/gin-gonic/gin"
)

var Movie struct {
	Title       string
	ReleaseYear int
	Rating      float32
}

func CreateMovie(c *gin.Context) {
	c.Bind(&Movie)
	movie := models.Movie{Title: Movie.Title, ReleaseYear: Movie.ReleaseYear, Rating: Movie.Rating}
	result := initializers.DB.Create(&movie)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func GetMovies(c *gin.Context) {
	var movies []models.Movie
	initializers.DB.Find(&movies)

	c.JSON(200, gin.H{
		"movies": movies,
	})
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	initializers.DB.First(&movie, id)

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&Movie)
	var movie models.Movie
	initializers.DB.First(&movie, id)

	initializers.DB.Model(&movie).Updates(models.Movie{
		Title:       Movie.Title,
		ReleaseYear: Movie.ReleaseYear,
		Rating:      Movie.Rating,
	})

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Movie{}, id)

	c.Status(200)
}
