package services

import (
	"example/go-api/models"
	"example/go-api/repositories"
)

type MovieService interface {
	CreateMovie(models.Movie) models.Movie
	GetMovies() []models.Movie
	GetMovie(id int) (models.Movie, error)
	UpdateMovie(movie models.Movie) models.Movie
	DeleteMovie(id int)
}

type movieService struct {
	movieRepostiory repositories.MovieRepostiory
}

func NewMovieService(repo repositories.MovieRepostiory) MovieService {
	return &movieService{
		movieRepostiory: repo,
	}
}

func (service *movieService) CreateMovie(movie models.Movie) models.Movie {
	service.movieRepostiory.CreateMovie(movie)
	return movie
}

func (service *movieService) GetMovies() []models.Movie {
	return service.movieRepostiory.GetMovies()
}

func (service *movieService) GetMovie(id int) (models.Movie, error) {
	return service.movieRepostiory.GetMovie(id)
}

func (service *movieService) UpdateMovie(movie models.Movie) models.Movie {
	return service.movieRepostiory.UpdateMovie(movie)
}

func (service *movieService) DeleteMovie(id int) {
	service.movieRepostiory.DeleteMovie(id)
}
