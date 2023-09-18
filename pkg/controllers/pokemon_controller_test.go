package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/execb5/pokedex/database"
	"github.com/execb5/pokedex/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestControllers_PokemonIndex(t *testing.T) {
	database.InitializeTestDB()
	db := database.GetDB()

	arbok := models.Pokemon{
		ID:             24,
		Identifier:     "arbok",
		SpeciesId:      24,
		Height:         35,
		Weight:         650,
		BaseExperience: 157,
		Order:          33,
		IsDefault:      true,
	}
	pikachu := models.Pokemon{
		ID:             25,
		Identifier:     "pikachu",
		SpeciesId:      25,
		Height:         40,
		Weight:         60,
		BaseExperience: 112,
		Order:          35,
		IsDefault:      true,
	}

	db.Create(&pikachu)
	db.Create(&arbok)

	t.Cleanup(func() {
		db.Delete(&arbok)
		db.Delete(&pikachu)
	})

	request, err := http.NewRequest("GET", "/pokemon", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PokemonIndex)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response []models.Pokemon
	if err := json.NewDecoder(responseRecorder.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	if !reflect.DeepEqual(response[0], arbok) {
		t.Errorf(
			"models.Pokemon(%d, %s, %d, %d, %d, %d, %d, %v)=%v; want %v",
			response[0].ID,
			response[0].Identifier,
			response[0].SpeciesId,
			response[0].Height,
			response[0].Weight,
			response[0].BaseExperience,
			response[0].Order,
			response[0].IsDefault,
			response[0],
			arbok,
		)
	}
	if !reflect.DeepEqual(response[1], pikachu) {
		t.Errorf(
			"models.Pokemon(%d, %s, %d, %d, %d, %d, %d, %v)=%v; want %v",
			response[1].ID,
			response[1].Identifier,
			response[1].SpeciesId,
			response[1].Height,
			response[1].Weight,
			response[1].BaseExperience,
			response[1].Order,
			response[1].IsDefault,
			response[1],
			pikachu,
		)
	}
}

func TestControllers_findAllPokemon(t *testing.T) {
	database.InitializeTestDB()
	db := database.GetDB()

	arbok := models.Pokemon{
		ID:             24,
		Identifier:     "arbok",
		SpeciesId:      24,
		Height:         35,
		Weight:         650,
		BaseExperience: 157,
		Order:          33,
		IsDefault:      true,
	}
	pikachu := models.Pokemon{
		ID:             25,
		Identifier:     "pikachu",
		SpeciesId:      25,
		Height:         40,
		Weight:         60,
		BaseExperience: 112,
		Order:          35,
		IsDefault:      true,
	}

	db.Create(&pikachu)
	db.Create(&arbok)

	t.Cleanup(func() {
		db.Delete(&arbok)
		db.Delete(&pikachu)
	})

	pokemons := findAllPokemon()

	assert.Equal(t, pokemons[0], arbok, "they are supposed to be equal")
	assert.Equal(t, pokemons[1], pikachu, "they are supposed to be equal")
}
