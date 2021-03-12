package router

import (
	"github.com/gorilla/mux"
	"src/api/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
