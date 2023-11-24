package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPriceHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/price", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(priceHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expectedErrorMsg := "crypto parameter is required"
	if rr.Body.String() != expectedErrorMsg+"\n" {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedErrorMsg)
	}

}
