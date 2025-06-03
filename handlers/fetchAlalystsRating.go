package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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
	if apiToken == "" {
		return errors.New("API_TOKEN no está definido en las variables de entorno")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creando la petición HTTP: %v\n", err)
		return fmt.Errorf("error creando la petición HTTP: %w", err)
	}
	req.Header.Set("Authorization", apiToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error al realizar la petición HTTP: %v\n", err)
		return fmt.Errorf("error al realizar la petición HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Printf("Respuesta inesperada del API externo: %s\n", string(bodyBytes))
		return fmt.Errorf("respuesta inesperada del API externo: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error leyendo el cuerpo de la respuesta: %v\n", err)
		return fmt.Errorf("error leyendo el cuerpo de la respuesta: %w", err)
	}

	var apiResp apiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		fmt.Printf("Error al parsear el JSON de la respuesta: %v\n", err)
		return fmt.Errorf("error al parsear el JSON de la respuesta: %w", err)
	}

	if len(apiResp.Items) > 0 {
		if err := database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&apiResp.Items).Error; err != nil {
			fmt.Printf("Error guardando los datos en la base de datos: %v\n", err)
			return fmt.Errorf("error guardando los datos en la base de datos: %w", err)
		}
	}
	return nil
}
