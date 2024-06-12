package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/shariq/golang-dev-logic-challenge-shariqcheema/controllers"
	"github.com/shariq/golang-dev-logic-challenge-shariqcheema/model"
)

type AnalysisResult struct {
	GraphData       []controllers.XYValue `json:"graph_data"`
	MaxProfit       float64                `json:"max_profit"`
	MaxLoss         float64                `json:"max_loss"`
	BreakEvenPoints []float64              `json:"break_even_points"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
		var contracts []model.OptionsContract

		if err := c.ShouldBindJSON(&contracts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		for _, contract := range contracts {
			if contract.Type != "Call" && contract.Type != "Put" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid option type"})
				return
			}
			if contract.LongShort != "long" && contract.LongShort != "short" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position type"})
				return
			}
		}

		analysisResult := controllers.AnalyzeOptionsContracts(contracts)

		c.JSON(http.StatusOK, analysisResult)
	})

	return router
}
