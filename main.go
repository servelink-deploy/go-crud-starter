package main

import (
	"go-crud-starter/config"
	"go-crud-starter/routes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, utilisation des variables d'environnement système")
	}

	if err := config.InitDatabase(); err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la base de données: %v", err)
	}
	defer config.CloseDatabase()

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(corsConfig))

	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	go func() {
		log.Printf("✓ Serveur démarré sur le port %s", port)
		log.Printf("✓ Health check: http://localhost:%s/health", port)
		log.Printf("✓ API: http://localhost:%s/api/users", port)
		
		if err := router.Run(":" + port); err != nil {
			log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Arrêt gracieux du serveur...")
	config.CloseDatabase()
	log.Println("✓ Serveur arrêté proprement")
}
