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
		movies.GET("/", movieHandler.GetMovies)
		movies.GET("/:id", movieHandler.GetMovieById)

		movies.Use(middleware.AuthRequired())
		{
			movies.POST("/", movieHandler.CreateMovie)
			movies.PUT("/:id", movieHandler.UpdateMovie)
			movies.DELETE("/:id", movieHandler.DeleteMovie)
		}
	}
}
