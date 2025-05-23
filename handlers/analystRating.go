package handlers

import (
	"net/http"
	"strconv"

	"github.com/ALEJABM0817/TGolang/database"
	"github.com/ALEJABM0817/TGolang/models"
	"github.com/gin-gonic/gin"
)

const pageSize = 10

func GetAnalystRatings(c *gin.Context) {
	pageStr := c.DefaultQuery("next_page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	var ratings []models.AnalystRating
	result := database.DB.Order("id").Limit(pageSize).Offset(offset).Find(&ratings)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(ratings) < pageSize {
		if err := FetchAndSaveAnalystRatingsUtil(pageStr); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch and save: " + err.Error()})
			return
		}
		ratings = nil
		result = database.DB.Order("id").Limit(pageSize).Offset(offset).Find(&ratings)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, ratings)
}
