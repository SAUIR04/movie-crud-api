package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"movie-api/internal/db"
	"movie-api/internal/routes"
)

func main() {
	// Дерекқорды инициализациялау
	db.InitDB()
	// Gin серверін бастау
	r := gin.Default()
	// Роуттарды орнату
	routes.SetupRoutes(r, db.DB)

	// Серверді іске қосу
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
