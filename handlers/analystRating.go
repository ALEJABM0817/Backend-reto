package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ALEJABM0817/TGolang/database"
	"github.com/ALEJABM0817/TGolang/models"
	"github.com/gin-gonic/gin"
)

const pageSize = 10

var validSortKeys = map[string]bool{
	"company":     true,
	"time":        true,
	"ticker":      true,
	"target_from": true,
	"target_to":   true,
	"action":      true,
	"brokerage":   true,
	"rating_from": true,
	"rating_to":   true,
	"id":          true,
}

func getOrderParams(c *gin.Context) (string, string) {
	sortKey := c.DefaultQuery("sortKey", "company")
	sortOrder := c.DefaultQuery("sortOrder", "desc")
	if !validSortKeys[sortKey] {
		sortKey = "company"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}
	return sortKey, sortOrder
}

func queryRatings(search, orderByUser string, offset int) ([]models.AnalystRating, error) {
	var ratings []models.AnalystRating
	db := database.DB.Order(orderByUser).Limit(pageSize).Offset(offset)
	if search != "" {
		like := "%" + search + "%"
		db = db.Where(
			"ticker ILIKE ? OR company ILIKE ? OR action ILIKE ? OR brokerage ILIKE ? OR rating_from ILIKE ? OR rating_to ILIKE ?",
			like, like, like, like, like, like,
		)
	}
	result := db.Find(&ratings)
	return ratings, result.Error
}

func GetAnalystRatings(c *gin.Context) {
	pageStr := c.DefaultQuery("next_page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	search := c.Query("search")
	sortKey, sortOrder := getOrderParams(c)
	orderByUser := fmt.Sprintf("%s %s", sortKey, sortOrder)

	ratings, err := queryRatings(search, orderByUser, offset)
	if err != nil {
		fmt.Printf("Error al consultar la base de datos: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al consultar los datos"})
		return
	}

	if len(ratings) < pageSize && search == "" {
		if err := FetchAndSaveAnalystRatingsUtil(pageStr); err != nil {
			fmt.Printf("Error al obtener y guardar datos externos: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener más datos externos"})
			return
		}
		ratings, err = queryRatings(search, orderByUser, offset)
		if err != nil {
			fmt.Printf("Error al consultar la base de datos tras fetch: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al consultar los datos después de actualizar"})
			return
		}
	}

	c.JSON(http.StatusOK, ratings)
}
