package repositories

import (
	"example/go-api/initializers"
	"example/go-api/models"
)

type ReviewRepository interface {
	GetReview(id int) models.Review
	GetAllMovieReviews(movieId int) []models.Review
	AddReview(models.Review) models.Review
	UpdateReview(models.Review) models.Review
	DeleteReview(id int) models.Review
}

type reviewRepository struct {
}

func NewReviewRepository() ReviewRepository {
	return reviewRepository{}
}

func (reviewRepository) AddReview(review models.Review) models.Review {
	initializers.DB.Create(&review)
	return review
}

func (reviewRepository) GetReview(id int) models.Review {
	var review models.Review
	err := initializers.DB.First(&review, id).Error
	if err != nil {
		// Handle error
	}

	return review
}

func (reviewRepository) GetAllMovieReviews(movieId int) []models.Review {
	var reviews []models.Review
	err := initializers.DB.Where("movie_id = ?", movieId).Find(&reviews).Error
	if err != nil {
		// Handle error
	}

	return reviews
}

func (reviewRepository) UpdateReview(review models.Review) models.Review {
	err := initializers.DB.Model(&review).Updates(models.Review{
		Rating: review.Rating,
		Review: review.Review,
	})

	if err != nil {
		// Handle error
	}
	return review
}

func (reviewRepository) DeleteReview(id int) models.Review {
	var review models.Review
	initializers.DB.Delete(&models.Review{}, id)
	return review
}
