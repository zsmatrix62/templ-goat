package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterStatic(r *mux.Router) {
	r.PathPrefix("/files/").Handler(
		http.StripPrefix("/files/",
			http.FileServer(http.Dir("web/downloads")),
		),
	)
}
