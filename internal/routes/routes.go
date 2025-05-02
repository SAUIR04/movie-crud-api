package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-api/internal/delivery"
	"movie-api/internal/middleware"
	"movie-api/internal/repository"
	"movie-api/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	delivery.SetAuthDB(db)

	r.POST("/register", delivery.Register)
	r.POST("/login", delivery.Login)

	movieRepo := repository.NewMovieRepository(db)
	movieService := services.NewMovieService(movieRepo)
	movieHandler := delivery.NewMovieHandler(movieService)

	movies := r.Group("/api/movies")
	{
		// GET запросы доступны всем
		movies.GET("/", movieHandler.GetMovies)       // Получение всех фильмов
		movies.GET("/:id", movieHandler.GetMovieById) // Получение фильма по ID

		// Применяем middleware для аутентификации
		movies.Use(middleware.AuthRequired())
		{
			// Только для администраторов
			movies.Use(middleware.AdminRequired())
			{
				movies.POST("/", movieHandler.CreateMovie)      // Создание фильма
				movies.PUT("/:id", movieHandler.UpdateMovie)    // Обновление фильма
				movies.DELETE("/:id", movieHandler.DeleteMovie) // Удаление фильма
			}
		}
	}
}
