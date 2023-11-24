package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sahilpatel17/week12/prices"
)

// GetPrice is the handler for the "price" endpoint
func GetPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	crypto := params["crypto"]

	price, err := prices.GetCryptoPrice(crypto)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching price for %s: %s", crypto, err), http.StatusInternalServerError)
		return
	}

	response := map[string]float64{crypto: price}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
