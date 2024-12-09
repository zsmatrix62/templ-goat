package bootstrap

import (
	"embed"
	"io/fs"
	"net/http"
	"github.com/zsmatrix62/templ-goat/app/http/middlewares"
	"github.com/zsmatrix62/templ-goat/pkg/route"
	"github.com/zsmatrix62/templ-goat/routes"

	"github.com/gorilla/mux"
)

func SetupRoutes(publicFS embed.FS) *mux.Router {
	route.Router.Use(middlewares.ForceHTMLMiddleware)
	route.Router.Use(middlewares.RemoveTrailingSlash)

	r1 := route.Router.NewRoute().Subrouter()
	routes.RegisterIndexRoutes(r1)

	sub, _ := fs.Sub(publicFS, "public")
	route.Router.PathPrefix("/").Handler(http.FileServer(http.FS(sub)))
	routes.RegisterStatic(route.Router)
	return route.Router
}
