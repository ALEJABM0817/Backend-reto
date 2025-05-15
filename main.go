package main

import (
	"github.com/ALEJABM0817/TGolang/Backend/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
