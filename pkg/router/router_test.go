package router

import (
	"reflect"
	"testing"

	"github.com/execb5/pokedex/pkg/controllers"
	"github.com/stretchr/testify/assert"
)

func TestRouter_RoutesShoulReturnCorrectControllerHandlers(t *testing.T) {
	routes := Routes()

	assert.True(t, reflect.ValueOf(routes["/"]).Pointer() == reflect.ValueOf(controllers.Index).Pointer(), "should be the same function")
	assert.True(t, reflect.ValueOf(routes["/pokemon"]).Pointer() == reflect.ValueOf(controllers.PokemonIndex).Pointer(), "should be the same function")
	assert.True(t, reflect.ValueOf(routes["/pokemon/"]).Pointer() == reflect.ValueOf(controllers.PokemonShow).Pointer(), "should be the same function")
}

func TestRouter_RoutesShoulContainOnlyCorrectRoutes(t *testing.T) {
	expectedPaths := []string{"/", "/pokemon", "/pokemon/"}

	pathsFound := make(map[string]bool)

	for key := range Routes() {
		pathsFound[key] = true
	}

	for _, expectedPath := range expectedPaths {
		if !pathsFound[expectedPath] {
			t.Errorf("router has unexpected key -> %v", expectedPath)
		}
	}

	if len(pathsFound) != len(expectedPaths) {
		paths := []string{}
		for key := range pathsFound {
			paths = append(paths, key)
		}
		t.Errorf("router has unexpected paths -> found: %v, expected: %v", paths, expectedPaths)
	}

}
