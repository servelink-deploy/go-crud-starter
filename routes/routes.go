package routes

import (
	"go-crud-starter/config"
	"go-crud-starter/handlers"
	"go-crud-starter/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		if err := config.DB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   "unhealthy",
				"error":    "Erreur de connexion à la base de données",
				"service":  "go-crud-api",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "healthy",
			"database": "connected",
			"service":  "go-crud-api",
		})
	})

	userHandler := handlers.NewUserHandler()

	api := router.Group("/api")
	api.Use(middleware.RateLimiter(100))
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetAllUsers)
			users.GET("/search", userHandler.SearchUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route non trouvée",
		})
	})
}
