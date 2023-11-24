package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/go-crypto-price-tracker/api"
)

func TestGetPrice(t *testing.T) {
	req, err := http.NewRequest("GET", "/price/bitcoin", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetPrice)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Add more assertions based on your requirements
}
