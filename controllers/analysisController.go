package controllers

import (
	"time"
	"math"
	"github.com/shariq/golang-dev-logic-challenge-shariqcheema/model"
)

type OptionsContract struct {
	Type          string    `json:"type"`
	StrikePrice   float64   `json:"strike_price"`
	Bid           float64   `json:"bid"`
	Ask           float64   `json:"ask"`
	ExpirationDate time.Time `json:"expiration_date"`
	LongShort     string    `json:"long_short"`
}

type AnalysisResponse struct {
	XYValues        []XYValue `json:"xy_values"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}

type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func AnalyzeOptionsContracts(contracts []model.OptionsContract) AnalysisResponse {
	xyValues := calculateXYValues(contracts)
	maxProfit := calculateMaxProfit(contracts)
	maxLoss := calculateMaxLoss(contracts)
	breakEvenPoints := calculateBreakEvenPoints(contracts)

	return AnalysisResponse{
		XYValues:        xyValues,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}
}

func calculateXYValues(contracts []model.OptionsContract) []XYValue {
	var xyValues []XYValue
	for price := 50.0; price <= 150.0; price += 1.0 {
		profitLoss := 0.0
		for _, contract := range contracts {
			if contract.Type == "Call" {
				if contract.LongShort == "long" {
					profitLoss += math.Max(0, price-contract.StrikePrice) - contract.Ask
				} else {
					profitLoss += contract.Bid - math.Max(0, price-contract.StrikePrice)
				}
			} else if contract.Type == "Put" {
				if contract.LongShort == "long" {
					profitLoss += math.Max(0, contract.StrikePrice-price) - contract.Ask
				} else {
					profitLoss += contract.Bid - math.Max(0, contract.StrikePrice-price)
				}
			}
		}
		xyValues = append(xyValues, XYValue{X: price, Y: profitLoss})
	}
	return xyValues
}

func calculateMaxProfit(contracts []model.OptionsContract) float64 {
	maxProfit := math.Inf(-1)
	for _, xy := range calculateXYValues(contracts) {
		if xy.Y > maxProfit {
			maxProfit = xy.Y
		}
	}
	return maxProfit
}

func calculateMaxLoss(contracts []model.OptionsContract) float64 {
	maxLoss := math.Inf(1)
	for _, xy := range calculateXYValues(contracts) {
		if xy.Y < maxLoss {
			maxLoss = xy.Y
		}
	}
	return maxLoss
}

func calculateBreakEvenPoints(contracts []model.OptionsContract) []float64 {
	breakEvenPoints := []float64{}
	xyValues := calculateXYValues(contracts)
	for i := 1; i < len(xyValues); i++ {
		if xyValues[i-1].Y*xyValues[i].Y <= 0 {
			breakEvenPoints = append(breakEvenPoints, xyValues[i].X)
		}
	}
	return breakEvenPoints
}
