package main

import (
	"github.com/ALEJABM0817/TGolang/database"
	"github.com/ALEJABM0817/TGolang/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/analyst-ratings", handlers.GetAnalystRatings)
	r.POST("/fetch-analyst-ratings", handlers.FetchAndSaveAnalystRatings)
	r.Run(":8080")
}
