package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/ALEJABM0817/TGolang/database"
	"github.com/ALEJABM0817/TGolang/models"
	"github.com/gin-gonic/gin"
)

func parseTarget(target string) float64 {
	re := regexp.MustCompile(`[0-9.]+`)
	val := re.FindString(target)
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Printf("Error al convertir target '%s' a float: %v\n", target, err)
		return 0
	}
	return f
}

type stat struct {
	Ticker    string
	BuyCount  int
	MaxTarget float64
	Company   string
}

func RecommendBestStock(c *gin.Context) {
	var ratings []models.AnalystRating
	if err := database.DB.Find(&ratings).Error; err != nil {
		fmt.Printf("Error al consultar la base de datos: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al consultar los datos"})
		return
	}

	stats := map[string]*stat{}

	for _, r := range ratings {
		ticker := r.Ticker
		if _, ok := stats[ticker]; !ok {
			stats[ticker] = &stat{Ticker: ticker, Company: r.Company}
		}
		if strings.EqualFold(r.RatingTo, "Buy") {
			stats[ticker].BuyCount++
			target := parseTarget(r.TargetTo)
			if target > stats[ticker].MaxTarget {
				stats[ticker].MaxTarget = target
			}
		}
	}

	var statList []*stat
	for _, s := range stats {
		statList = append(statList, s)
	}

	if len(statList) == 0 {
		c.JSON(http.StatusOK, gin.H{"recommendation": "No hay datos disponibles"})
		return
	}

	sort.Slice(statList, func(i, j int) bool {
		if statList[i].BuyCount == statList[j].BuyCount {
			return statList[i].MaxTarget > statList[j].MaxTarget
		}
		return statList[i].BuyCount > statList[j].BuyCount
	})

	best := statList[0]
	c.JSON(http.StatusOK, gin.H{
		"recommendation": best.Ticker,
		"company":        best.Company,
		"buy_count":      best.BuyCount,
		"max_target_to":  best.MaxTarget,
	})
}
