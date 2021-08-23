package routes

import (
	"github.com/galloaleonardo/api-marvel-heroes/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/heroes", controllers.Heroes)
	r.HandleFunc("/heroes/marvel", controllers.MarvelHeroes)
	r.HandleFunc("/heroes/dc", controllers.DCHeroes)

	return r
}
