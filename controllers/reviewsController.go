package controllers

import (
	"example/go-api/models"
	"example/go-api/services"

	"github.com/gin-gonic/gin"
)

type ReviewsController interface {
	AddReview(ctx *gin.Context)
	// UpdateReview(ctx *gin.Context)
	// ReadReviews(ctx *gin.Context)
	// DeleteReview(ctx *gin.Context)
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

	ctx.Param("id")

}
