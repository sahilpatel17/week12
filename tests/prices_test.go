package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yourusername/go-crypto-price-tracker/prices"
)

func TestGetCryptoPrice(t *testing.T) {
	price, err := prices.GetCryptoPrice("bitcoin")
	assert.NoError(t, err)
	assert.NotEqual(t, 0.0, price)
	// Add more assertions based on your requirements
}
