package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/galloaleonardo/api-marvel-heroes/services"
)

func Heroes(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	heroes := services.Heroes(params)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(heroes)
}

func MarvelHeroes(w http.ResponseWriter, r *http.Request) {
	heroes := services.MarvelHeroes()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(heroes)
}

func DCHeroes(w http.ResponseWriter, r *http.Request) {
	heroes := services.DCHeroes()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(heroes)
}
