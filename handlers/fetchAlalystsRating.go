package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ALEJABM0817/TGolang/database"
	"github.com/ALEJABM0817/TGolang/models"
	"gorm.io/gorm/clause"
)

const apiURL = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"

type apiResponse struct {
	Items    []models.AnalystRating `json:"items"`
	NextPage string                 `json:"next_page"`
}

func FetchAndSaveAnalystRatingsUtil(nextPage string) error {
	url := apiURL
	if nextPage != "" {
		url += "?next_page=" + nextPage
	}

	apiToken := os.Getenv("API_TOKEN")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResp apiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return err
	}

	if len(apiResp.Items) > 0 {
		if err := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&apiResp.Items).Error; err != nil {
			return err
		}
	}
	return nil
}
