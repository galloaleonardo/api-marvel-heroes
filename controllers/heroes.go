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
