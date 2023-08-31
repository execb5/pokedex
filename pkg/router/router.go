package router

import (
	"net/http"

	"github.com/execb5/pokedex/pkg/controllers"
)

func Routes() map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"/":        controllers.Index,
		"/pokemon": controllers.PokemonShow,
	}
}
