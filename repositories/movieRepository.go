package repositories

import (
	"example/go-api/initializers"
	"example/go-api/models"
)

type MovieRepostiory interface {
	CreateMovie(models.Movie)
	GetMovies() []models.Movie
	UpdateMovie(models.Movie) models.Movie
	DeleteMovie(id int)
	GetMovie(id int) (models.Movie, error)
}

type movieRepostiory struct {
}

func NewMovieRepository() MovieRepostiory {
	return movieRepostiory{}
}

func (movieRepostiory) CreateMovie(movie models.Movie) {
	initializers.DB.Create(&movie)
}

func (movieRepostiory) GetMovies() []models.Movie {
	var movies []models.Movie
	initializers.DB.Find(&movies)
	return movies
}

func (movieRepostiory) GetMovie(id int) (models.Movie, error) {
	var movie models.Movie
	if err := initializers.DB.First(&movie, id).Error; err != nil {
		return movie, err
	}
	return movie, nil
}

func (movieRepostiory) UpdateMovie(movie models.Movie) models.Movie {
	initializers.DB.Model(&movie).Updates(models.Movie{
		Title:       movie.Title,
		ReleaseYear: movie.ReleaseYear,
		Plot:        movie.Plot,
	})
	return movie
}

func (movieRepostiory) DeleteMovie(id int) {
	initializers.DB.Delete(&models.Movie{}, id)
}
