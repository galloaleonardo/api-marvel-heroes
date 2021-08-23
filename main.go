package main

import (
	"net/http"

	"github.com/galloaleonardo/api-marvel-heroes/routes"
)

func main() {
	r := routes.GetRoutes()

	http.ListenAndServe(":8001", r)
}
