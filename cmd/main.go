package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"movie-api/internal/db"
	"movie-api/internal/routes"
)

func main() {
	// Дерекқор параметрлерін .env файлы арқылы жүктеу

	// Дерекқорды инициализациялау
	db.InitDB()

	// Автоматты миграция
	//err := db.DB.AutoMigrate(&models.Movie{}, &models.User{})
	//if err != nil {
	//	log.Fatal("Error migrating to the DB:", err)
	//}

	// Gin серверін бастау
	r := gin.Default()

	// Роуттарды орнату
	routes.SetupRoutes(r, db.DB)

	// Серверді іске қосу
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
