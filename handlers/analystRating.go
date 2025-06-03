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

func GetAnalystRatings(c *gin.Context) {
	pageStr := c.DefaultQuery("next_page", "1")
	page, err := strconv.Atoi(pageStr)
	search := c.Query("search")

	if err != nil || page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	var ratings []models.AnalystRating

	if search != "" {
		like := "%" + search + "%"
		result := database.DB.
			Where(
				"ticker ILIKE ? OR company ILIKE ? OR action ILIKE ? OR brokerage ILIKE ? OR rating_from ILIKE ? OR rating_to ILIKE ?",
				like, like, like, like, like, like,
			).
			Order("id").
			Find(&ratings)
		if result.Error != nil {
			fmt.Printf("Error al obtener los datos en la busqueda de la base de datos: %v\n", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, ratings)
		return
	}

	result := database.DB.Order("id").Limit(pageSize).Offset(offset).Find(&ratings)

	if result.Error != nil {
		fmt.Printf("Error al consultar la base de datos: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al consultar los datos"})
		return
	}

	if len(ratings) < pageSize && search == "" {
		if err := FetchAndSaveAnalystRatingsUtil(pageStr); err != nil {
			fmt.Printf("Error al obtener y guardar datos externos: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener más datos externos"})
			return
		}
		ratings = nil
		result = database.DB.Order("id").Limit(pageSize).Offset(offset).Find(&ratings)
		if result.Error != nil {
			fmt.Printf("Error al consultar la base de datos tras fetch: %v\n", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al consultar los datos después de actualizar"})
			return
		}
	}

	c.JSON(http.StatusOK, ratings)
}
