package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-api/internal/delivery"
	"movie-api/internal/repository"
	"movie-api/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialize repositories
	movieRepo := repository.NewMovieRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	movieService := services.NewMovieService(movieRepo)
	userService := services.NewUserService(userRepo)

	// Initialize handlers
	movieHandler := delivery.NewMovieHandler(movieService)
	userHandler := delivery.NewUserHandler(userService)

	// Movie routes
	movies := r.Group("api/movies")
	{
		movies.GET("/", movieHandler.GetMovies)
		movies.POST("/", movieHandler.CreateMovie)
		movies.GET("/:id", movieHandler.GetMovieById)
		movies.PUT("/:id", movieHandler.UpdateMovie)
		movies.DELETE("/:id", movieHandler.DeleteMovie)
	}

	// User routes
	users := r.Group("api/users")
	{
		users.GET("/", userHandler.GetUsers)
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUserById)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}
