package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahilpatel17/week12/prices"
)

func main() {
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/prices", getAllPrices).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getAllPrices(w http.ResponseWriter, r *http.Request) {
	cryptos := []string{"bitcoin", "ethereum", "tether"}

	pricesMap, err := prices.GetCryptoPrices(cryptos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching prices: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pricesMap)
}
