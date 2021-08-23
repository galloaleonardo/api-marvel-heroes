package routes

import (
	"github.com/galloaleonardo/api-marvel-heroes/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.GetAllHeroes)

	return r
}