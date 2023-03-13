package repositories

import (
	"example/go-api/initializers"
	"example/go-api/models"
)

type ReviewRepository interface {
	GetReview(id int) (models.Review, error)
	GetAllMovieReviews(movieId int) ([]models.Review, error)
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

func (reviewRepository) GetReview(id int) (models.Review, error) {
	var review models.Review
	if err := initializers.DB.First(&review, id).Error; err != nil {
		return review, err
	}

	return review, nil
}

func (reviewRepository) GetAllMovieReviews(movieId int) ([]models.Review, error) {
	var reviews []models.Review
	if err := initializers.DB.Where("movie_id = ?", movieId).Find(&reviews).Error; err != nil {
		return reviews, err
	}
	return reviews, nil
}

func (reviewRepository) UpdateReview(review models.Review) models.Review {
	initializers.DB.Model(&review).Updates(models.Review{
		Rating: review.Rating,
		Review: review.Review,
	})
	return review
}

func (reviewRepository) DeleteReview(id int) models.Review {
	var review models.Review
	initializers.DB.Delete(&models.Review{}, id)
	return review
}
