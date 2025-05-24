package main

import (
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

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	database.Connect()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/analyst-ratings", handlers.GetAnalystRatings)
	r.GET("/recommendation", handlers.RecommendBestStock)

	r.Run(":8080")
}
