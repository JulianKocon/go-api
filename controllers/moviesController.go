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

type MovieController interface {
	GetMovies(ctx *gin.Context)
	CreateMovie(ctx *gin.Context)
	GetMovie(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
}

type controller struct {
	service services.MovieService
}

func New(service services.MovieService) MovieController {
	return &controller{
		service: service,
	}
}

func (c *controller) CreateMovie(ctx *gin.Context) {
	var movie models.Movie
	err := ctx.ShouldBindJSON(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Param("id")
	c.service.CreateMovie(movie)
	ctx.JSON(http.StatusCreated, &movie)
}

func (c *controller) GetMovies(ctx *gin.Context) {
	movies := c.service.GetMovies()
	ctx.JSON(http.StatusOK, &movies)
}

func (c *controller) GetMovie(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie := c.service.GetMovie(id)
	ctx.JSON(http.StatusOK, &movie)
}

func (c *controller) UpdateMovie(ctx *gin.Context) {
	var movie models.Movie
	err := ctx.BindJSON(&movie)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie.ID = uint(id)
	c.service.UpdateMovie(movie)
	ctx.JSON(http.StatusOK, &movie)
}

func (c *controller) DeleteMovie(ctx *gin.Context) {
	var movie models.Movie
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie.ID = uint(id)
	c.service.DeleteMovie(id)
	ctx.JSON(http.StatusOK, &movie)
}
