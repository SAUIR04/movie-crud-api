package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"movie-api/internal/delivery"
	"movie-api/internal/models"
	"movie-api/internal/routes"
)

func main() {
	dsn := "postgres://postgres:2004@localhost:5000/mydatabase?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	// Автоматты миграция
	db.AutoMigrate(&models.Movie{}, &models.User{})

	// Auth DB setup
	delivery.SetAuthDB(db)

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
