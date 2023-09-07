package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestControllers_Index(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	if err := json.NewDecoder(responseRecorder.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	if message, ok := response["message"]; !ok || message != "Hello. I'm the Pokedex server" {
		t.Errorf("Unexpected 'message' field in JSON response: got %v want 'Hello. I'm the Pokedex server'", message)
	}
}
