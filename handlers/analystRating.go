package handlers

import (
	"net/http"

	"github.com/ALEJABM0817/TGolang/database"
	"github.com/ALEJABM0817/TGolang/models"
	"github.com/gin-gonic/gin"
)

func GetAnalystRatings(c *gin.Context) {
	var ratings []models.AnalystRating
	if err := database.DB.Find(&ratings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ratings)
}
