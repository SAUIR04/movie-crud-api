package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie-api/internal/models"
	"movie-api/internal/services"
	"net/http"
)

type MovieHandler struct {
	service *services.MovieService
}

func NewMovieHandler(service *services.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

// Получение списка всех фильмов
func (h *MovieHandler) GetMovies(c *gin.Context) {
	movies, err := h.service.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения фильмов"})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// Получение фильма по ID
func (h *MovieHandler) GetMovieById(c *gin.Context) {
	id := c.Param("id")
	// Преобразуем id в тип uint
	var movieId uint
	if _, err := fmt.Sscanf(id, "%d", &movieId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	movie, err := h.service.GetMovieById(movieId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Фильм не найден"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// Создание нового фильма
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	err := h.service.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания фильма"})
		return
	}
	c.JSON(http.StatusCreated, movie)
}

// Обновление информации о фильме
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	// Преобразуем id в тип uint
	var movieId uint
	if _, err := fmt.Sscanf(id, "%d", &movieId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err := h.service.UpdateMovie(movieId, movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления фильма"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// Удаление фильма
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	// Преобразуем id в тип uint
	var movieId uint
	if _, err := fmt.Sscanf(id, "%d", &movieId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err := h.service.DeleteMovie(movieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления фильма"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
