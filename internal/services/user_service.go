package services

import (
	"errors"
	"movie-api/internal/models"
	"movie-api/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) CreateUser(user models.User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("все поля обязательны")
	}
	return s.repo.SaveUser(user)
}

func (s *UserService) UpdateUser(id uint, user models.User) error {
	return s.repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
