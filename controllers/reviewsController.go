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
	err := ctx.ShouldBindJSON(&review)
	ReturnIfError(ctx, err)
	movieId, err := strconv.Atoi(ctx.Param("id"))
	ReturnIfError(ctx, err)
	review.MovieID = uint64(movieId);
	ctx.Param("id")
	c.service.AddReview(review)
	ctx.JSON(http.StatusCreated, &review)
}

func (c *reviewsController) GetReviews(ctx *gin.Context) {
	movieId, err := strconv.Atoi(ctx.Param("id"))
	ReturnIfError(ctx, err)
	reviews := c.service.GetAllMovieReviews(movieId)
	ctx.JSON(http.StatusOK, &reviews)
}

func (c *reviewsController) GetReviewById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	ReturnIfError(ctx, err)
	review := c.service.GetReview(id)
	ctx.JSON(http.StatusOK, &review)
}

func (c *reviewsController) UpdateReview(ctx *gin.Context) {
	var review models.Review
	err := ctx.ShouldBindJSON(&review)
	ReturnIfError(ctx, err)
	id, err := strconv.Atoi(ctx.Param("id"))
	ReturnIfError(ctx, err)
	review.ID = uint(id)
	c.service.UpdateReview(review)
	ctx.Status(http.StatusOK)
}

func (c *reviewsController) DeleteReview(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	ReturnIfError(ctx, err)
	c.service.DeleteReview(id)
	ctx.Status(http.StatusOK)
}
