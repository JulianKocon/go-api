package repositories

import (
	"example/go-api/initializers"
	"example/go-api/models"
)

type MovieRepostiory interface {
	CreateMovie(models.Movie)
	GetMovies() []models.Movie
	UpdateMovie(models.Movie) models.Movie
	DeleteMovie(id int) models.Movie
	GetMovie(id int) models.Movie
	GetMovieRating(movieId int) (float32, error)
}

type movieRepostiory struct {
}

func NewMovieRepository() MovieRepostiory {
	return movieRepostiory{}
}

func (movieRepostiory) GetMovieRating(movieId int) (float32, error) {
	var result float32
	row := initializers.DB.Table("reviews").Select("avg(rating)").Where("movie_id = ?", movieId).Row()
	if err :=row.Scan(&result); err != nil{
		return 0, err
	}
	return result, nil
}

func (movieRepostiory) CreateMovie(movie models.Movie) {
	initializers.DB.Create(&movie)
}

func (movieRepostiory) GetMovies() []models.Movie {
	var movies []models.Movie
	initializers.DB.Find(&movies)
	return movies
}

func (movieRepostiory) GetMovie(id int) models.Movie {
	var movie models.Movie
	initializers.DB.First(&movie, id)
	return movie
}

func (movieRepostiory) UpdateMovie(movie models.Movie) models.Movie {
	initializers.DB.First(&movie, movie.ID)
	initializers.DB.Model(&movie).Updates(models.Movie{
		Title:       movie.Title,
		ReleaseYear: movie.ReleaseYear,
		Plot:        movie.Plot,
	})
	return movie
}

func (movieRepostiory) DeleteMovie(id int) models.Movie {
	var movie models.Movie
	initializers.DB.Delete(&models.Movie{}, id)
	return movie
}
