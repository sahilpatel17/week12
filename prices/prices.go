package prices

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://api.coingecko.com/api/v3/simple/price"

// GetCryptoPrice fetches the current price of the specified cryptocurrency
func GetCryptoPrice(crypto string) (float64, error) {
	url := fmt.Sprintf("%s?ids=%s&vs_currencies=cad", apiURL, crypto)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	price, ok := result[crypto]["cad"]
	if !ok {
		return 0, fmt.Errorf("price not found for %s", crypto)
	}

	return price, nil
}
