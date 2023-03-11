package controllers

import (
	"example/go-api/models"
	"example/go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewsController interface {
	AddReview(ctx *gin.Context)
	GetReviews(ctx *gin.Context)
	GetReviewById(ctx *gin.Context)
	UpdateReview(ctx *gin.Context)
	DeleteReview(ctx *gin.Context)
}

type reviewsController struct {
	service services.ReviewService
}

func NewReviewController(service services.ReviewService) ReviewsController {
	return &reviewsController{
		service: service,
	}
}

func (c *reviewsController) AddReview(ctx *gin.Context) {
	var review models.Review
	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	movieId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	review.MovieID = uint64(movieId)
	ctx.Param("id")
	c.service.AddReview(review)
	ctx.JSON(http.StatusCreated, &review)
}

func (c *reviewsController) GetReviews(ctx *gin.Context) {
	movieId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	reviews, err := c.service.GetAllMovieReviews(movieId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, &reviews)
}

func (c *reviewsController) GetReviewById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	review, err := c.service.GetReview(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, &review)
}

func (c *reviewsController) UpdateReview(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	review, err := c.service.GetReview(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	review.ID = uint(id)
	c.service.UpdateReview(review)
	ctx.JSON(http.StatusOK, &review)
}

func (c *reviewsController) DeleteReview(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.service.DeleteReview(id)
	ctx.Status(http.StatusOK)
}
