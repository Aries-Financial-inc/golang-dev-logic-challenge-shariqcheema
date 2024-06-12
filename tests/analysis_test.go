package tests

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"encoding/json"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/shariq/golang-dev-logic-challenge-shariqcheema/routes"
	"github.com/shariq/golang-dev-logic-challenge-shariqcheema/model"
)

func TestOptionsContractModelValidation(t *testing.T) {
	contract := model.OptionsContract{
		Type: "Call",
		StrikePrice: 100,
		Bid: 10.05,
		Ask: 12.04,
		LongShort: "long",
		ExpirationDate: time.Now(),
	}

	assert.Equal(t, "Call", contract.Type)
	assert.Equal(t, 100.0, contract.StrikePrice)
}

func TestAnalysisEndpoint(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/analyze", strings.NewReader(`[{
		"type": "Call",
		"strike_price": 100,
		"bid": 10.05,
		"ask": 12.04,
		"long_short": "long",
		"expiration_date": "2025-12-17T00:00:00Z"
	}]`))

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Contains(t, response, "xy_values")
	assert.Contains(t, response, "max_profit")
	assert.Contains(t, response, "max_loss")
	assert.Contains(t, response, "break_even_points")
}

func TestIntegration(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/analyze", strings.NewReader(`[{
		"type": "Call",
		"strike_price": 100,
		"bid": 10.05,
		"ask": 12.04,
		"long_short": "long",
		"expiration_date": "2025-12-17T00:00:00Z"
	}]`))

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Contains(t, response, "xy_values")
	assert.Contains(t, response, "max_profit")
	assert.Contains(t, response, "max_loss")
	assert.Contains(t, response, "break_even_points")
}

func TestMultipleContracts(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/analyze", strings.NewReader(`[{
		"type": "Call",
		"strike_price": 100,
		"bid": 10.05,
		"ask": 12.04,
		"long_short": "long",
		"expiration_date": "2025-12-17T00:00:00Z"
	}, {
		"type": "Put",
		"strike_price": 105,
		"bid": 16,
		"ask": 18,
		"long_short": "long",
		"expiration_date": "2025-12-17T00:00:00Z"
	}]`))

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Contains(t, response, "xy_values")
	assert.Contains(t, response, "max_profit")
	assert.Contains(t, response, "max_loss")
	assert.Contains(t, response, "break_even_points")
}

func TestInvalidData(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/analyze", strings.NewReader(`[{
		"type": "InvalidType",
		"strike_price": 100,
		"bid": 10.05,
		"ask": 12.04,
		"long_short": "long",
		"expiration_date": "2025-12-17T00:00:00Z"
	}]`))

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMixedOptionTypes(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/analyze", strings.NewReader(`[{
		"type": "Call",
		"strike_price": 100,
		"bid": 10.05,
		"ask": 12.04,
		"long_short": "long",
		"expiration_date": "2025-12-17T00:00:00Z"
	}, {
		"type": "Put",
		"strike_price": 105,
		"bid": 16,
		"ask": 18,
		"long_short": "short",
		"expiration_date": "2025-12-17T00:00:00Z"
	}, {
		"type": "Call",
		"strike_price": 110,
		"bid": 20.10,
		"ask": 22.05,
		"long_short": "long",
		"expiration_date": "2025-12-17T00:00:00Z"
	}, {
		"type": "Put",
		"strike_price": 95,
		"bid": 14,
		"ask": 15.50,
		"long_short": "short",
		"expiration_date": "2025-12-17T00:00:00Z"
	}]`))

	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Contains(t, response, "xy_values")
	assert.Contains(t, response, "max_profit")
	assert.Contains(t, response, "max_loss")
	assert.Contains(t, response, "break_even_points")
}
