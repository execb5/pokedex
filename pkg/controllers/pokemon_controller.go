package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/execb5/pokedex/database"
	"github.com/execb5/pokedex/pkg/models"
)

func PokemonShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(segments[2])

	db := database.GetDB()

	var pokemon models.Pokemon
	if err := db.First(&pokemon, id).Error; err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(pokemon)
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

func PokemonIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pokemons := findAllPokemon()

	jsonData, err := json.Marshal(pokemons)
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

func findAllPokemon() []models.Pokemon {
	var pokemons []models.Pokemon
	db := database.GetDB()
	db.Find(&pokemons)

	return pokemons
}
