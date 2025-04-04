package services

import (
	"errors"
	"movie-api/internal/models"
	"movie-api/internal/repository"
)

type MovieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.repo.GetAllMovies()
}

func (s *MovieService) GetMovieById(id uint) (*models.Movie, error) {
	return s.repo.GetMovieById(id)
}

func (s *MovieService) CreateMovie(movie models.Movie) error {
	if movie.Title == "" || movie.Description == "" {
		return errors.New("название и описание не могут быть пустыми")
	}
	return s.repo.SaveMovie(movie)
}

func (s *MovieService) UpdateMovie(id uint, movie models.Movie) error {
	return s.repo.UpdateMovie(id, movie)
}

func (s *MovieService) DeleteMovie(id uint) error {
	return s.repo.DeleteMovie(id)
}
