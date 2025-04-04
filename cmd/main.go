package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"movie-api/internal/models"
	"movie-api/internal/routes"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open("postgres://postgres:2004@localhost:5000/mydatabase?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err) // Логируем ошибку, если подключение не удалось
		return
	}

	// Auto-migrate models
	err = db.AutoMigrate(&models.Movie{}, &models.User{})
	if err != nil {
		log.Fatal("Error on migrating to the DB:", err) // Логируем ошибку миграции
		return
	}

	// Initialize Gin router
	r := gin.Default()

	// Set up routes
	routes.SetupRoutes(r, db)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err) // Логируем ошибку запуска сервера
		return
	}
}
