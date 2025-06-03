package main

import (
	"log"
	"os"
	"time"

	"github.com/ALEJABM0817/TGolang/database"
	"github.com/ALEJABM0817/TGolang/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
		log.Println("FRONTEND_URL no definido, usando valor por defecto:", frontendURL)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	database.Connect()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/analyst-ratings", handlers.GetAnalystRatings)
	r.GET("/recommendation", handlers.RecommendBestStock)

	if err := r.Run(":8082"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
