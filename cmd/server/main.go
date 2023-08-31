package main

import (
	"fmt"
	"net/http"

	"github.com/execb5/pokedex/pkg/router"
)

func main() {
	for path, controllerFunc := range router.Routes() {
		http.HandleFunc(path, controllerFunc)
	}

	port := 8080
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
