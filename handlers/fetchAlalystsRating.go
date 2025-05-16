package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ALEJABM0817/TGolang/database"
	"github.com/ALEJABM0817/TGolang/models"
	"github.com/gin-gonic/gin"
)

const apiURL = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"

type apiResponse struct {
	Items    []models.AnalystRating `json:"items"`
	NextPage string                 `json:"next_page"`
}

func FetchAndSaveAnalystRatings(c *gin.Context) {
	nextPage := c.Query("next_page")
	url := apiURL
	if nextPage != "" {
		url += "?next_page=" + nextPage
	}

	apiToken := os.Getenv("API_TOKEN")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Authorization", apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from external API"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var apiResp apiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		return
	}

	for _, item := range apiResp.Items {
		if err := database.DB.Create(&item).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert item: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"inserted":  len(apiResp.Items),
		"next_page": apiResp.NextPage,
	})
}
