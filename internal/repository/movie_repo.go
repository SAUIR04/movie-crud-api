package repository

import (
	"gorm.io/gorm"
	"movie-api/internal/models"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return MovieRepository{db: db}
}

func (r *MovieRepository) SaveMovie(movie models.Movie) error {
	return r.db.Create(&movie).Error
}

func (r *MovieRepository) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) GetMovieById(id uint) (*models.Movie, error) {
	var movie models.Movie
	err := r.db.First(&movie, id).Error
	return &movie, err
}

func (r *MovieRepository) UpdateMovie(id uint, movie models.Movie) error {
	return r.db.Model(&models.Movie{}).Where("id = ?", id).Updates(movie).Error
}

func (r *MovieRepository) DeleteMovie(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}
