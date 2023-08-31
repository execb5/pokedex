package controllers

import (
	"encoding/json"
	"net/http"
)

func PokemonShow(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello. I'm the Pokemon show route",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
