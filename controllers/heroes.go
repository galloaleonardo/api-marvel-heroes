package controllers

import (
	"net/http"
)

func GetAllHeroes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing"))
}
