package handlers

import (
	"context"
	"net/http"

	"github.com/zsmatrix62/templ-goat/templates/pages"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	_ = pages.IndexPage().Render(context.Background(), w)
}
