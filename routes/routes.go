package routes

import (
	"net/http"
	"github.com/zsmatrix62/templ-goat/app/handlers"

	"github.com/gorilla/mux"
)

func RegisterIndexRoutes(r *mux.Router) {
	getR := r.NewRoute().Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/", handlers.IndexPage).Name("index")
}
