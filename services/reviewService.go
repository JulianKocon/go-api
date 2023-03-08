package services

import (
	"example/go-api/models"
	"example/go-api/repositories"
)

type ReviewService interface {
	GetReview(id int) models.Review
	GetAllMovieReviews(movieId int) []models.Review
	AddReview(models.Review) models.Review
	UpdateReview(models.Review) models.Review
	DeleteReview(id int) models.Review
}

type reviewService struct {
	reviewService repositories.ReviewRepository
}

func NewReviewService(repo repositories.ReviewRepository) ReviewService {
	return &reviewService{
		reviewService: repo,
	}
}

func (service *reviewService) GetReview(id int) models.Review {
	return service.reviewService.GetReview(id)
}

func (service *reviewService) GetAllMovieReviews(id int) []models.Review {
	return service.reviewService.GetAllMovieReviews(id)
}

func (service *reviewService) AddReview(review models.Review) models.Review {
	return service.reviewService.AddReview(review)
}

func (service *reviewService) UpdateReview(review models.Review) models.Review {
	return service.reviewService.UpdateReview(review)
}

func (service *reviewService) DeleteReview(id int) models.Review {
	return service.reviewService.DeleteReview(id)
}
