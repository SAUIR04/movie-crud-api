package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie-api/internal/models"
	"movie-api/internal/services"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Получение списка всех пользователей
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения пользователей"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Получение пользователя по ID
func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	// Преобразуем id в тип uint
	var userId uint
	if _, err := fmt.Sscanf(id, "%d", &userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}

	user, err := h.service.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Создание нового пользователя
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// Обновление информации о пользователе
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	// Преобразуем id в тип uint
	var userId uint
	if _, err := fmt.Sscanf(id, "%d", &userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err := h.service.UpdateUser(userId, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления пользователя"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Удаление пользователя
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// Преобразуем id в тип uint
	var userId uint
	if _, err := fmt.Sscanf(id, "%d", &userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID"})
		return
	}
	err := h.service.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления пользователя"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
