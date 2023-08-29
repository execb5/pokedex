package models

import "testing"

func TestModels_PokemonTableNameShouldReturnPokemon(t *testing.T) {
	expected := "pokemon"
	var pokemon Pokemon
	tableName := pokemon.TableName()
	if tableName != expected {
		t.Errorf("Table name expected to be %s, got: %s", expected, tableName)
	}
}
