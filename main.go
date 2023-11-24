package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CryptoPrice represents the JSON response structure from the external API.
type CryptoPrice struct {
	Bitcoin  map[string]float64 `json:"bitcoin"`
	Ethereum map[string]float64 `json:"ethereum"`
	Tether   map[string]float64 `json:"tether"`
}

// getPrice fetches the current price of Bitcoin, Ethereum, and Tether from the external API.
func getPrice() (CryptoPrice, error) {
	// You can replace the API URL with the one you selected.
	apiURL := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,tether&vs_currencies=cad"

	// Make an HTTP request to the external API.
	response, err := http.Get(apiURL)
	if err != nil {
		return CryptoPrice{}, err
	}
	defer response.Body.Close()

	// Decode the JSON response.
	var cryptoPrices CryptoPrice
	err = json.NewDecoder(response.Body).Decode(&cryptoPrices)
	if err != nil {
		return CryptoPrice{}, err
	}

	return cryptoPrices, nil
}

// priceHandler is the HTTP handler for the "price" endpoint.
func priceHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch the current prices for Bitcoin, Ethereum, and Tether.
	cryptoPrices, err := getPrice()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with individual JSON payloads for each coin.
	w.Header().Set("Content-Type", "application/json")

	// Bitcoin JSON response
	bitcoinJSON, err := json.Marshal(map[string]interface{}{
		"bitcoin": cryptoPrices.Bitcoin["cad"],
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", bitcoinJSON)

	// Ethereum JSON response
	ethereumJSON, err := json.Marshal(map[string]interface{}{
		"ethereum": cryptoPrices.Ethereum["cad"],
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", ethereumJSON)

	// Tether JSON response
	tetherJSON, err := json.Marshal(map[string]interface{}{
		"tether": cryptoPrices.Tether["cad"],
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", tetherJSON)
}

func main() {
	// Register the priceHandler function for the "/price" endpoint.
	http.HandleFunc("/price", priceHandler)

	// Start the HTTP server on port 8080.
	fmt.Println("Server listening on :3000")
	http.ListenAndServe(":3000", nil)
}
