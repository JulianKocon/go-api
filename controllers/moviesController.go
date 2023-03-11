package controllers

import (
	"example/go-api/models"
	"example/go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Movie struct {
	Title       string
	ReleaseYear int
	Rating      float32
}

type MoviesController interface {
	GetMovies(ctx *gin.Context)
	CreateMovie(ctx *gin.Context)
	GetMovie(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
	GetRating(ctx *gin.Context)
}

type moviesController struct {
	service services.MovieService
}

func NewMovieController(service services.MovieService) MoviesController {
	return &moviesController{
		service: service,
	}
}

func (c *moviesController) CreateMovie(ctx *gin.Context) {
	var movie models.Movie
	err := ctx.ShouldBindJSON(&movie)
	ReturnIfError(ctx, err)

	c.service.CreateMovie(movie)
	ctx.JSON(http.StatusCreated, &movie)
}

func (c *moviesController) GetMovies(ctx *gin.Context) {
	movies := c.service.GetMovies()
	ctx.JSON(http.StatusOK, &movies)
}

func (c *moviesController) GetMovie(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	ReturnIfError(ctx, err)

	movie := c.service.GetMovie(id)
	ctx.JSON(http.StatusOK, &movie)
}

func (c *moviesController) UpdateMovie(ctx *gin.Context) {
	var movie models.Movie
	err := ctx.BindJSON(&movie)
	ReturnIfError(ctx, err)

	id, err := strconv.Atoi(ctx.Param("id"))
	ReturnIfError(ctx, err)

	movie.ID = uint(id)
	c.service.UpdateMovie(movie)
	ctx.JSON(http.StatusOK, &movie)
}

func (c *moviesController) DeleteMovie(ctx *gin.Context) {
	var movie models.Movie
	id, err := strconv.Atoi(ctx.Param("id"))

	ReturnIfError(ctx, err)
	movie.ID = uint(id)
	c.service.DeleteMovie(id)
	ctx.JSON(http.StatusOK, &movie)
}

func (c *moviesController) GetRating(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movieRating, err := c.service.GetMovieRating(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &movieRating)
}

func ReturnIfError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
