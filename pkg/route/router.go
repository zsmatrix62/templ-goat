package route

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()
}

// RouteName2URL 通过路由名称来获取 URL
func RouteName2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}

	return url.String()
}

func Redirect(
	w http.ResponseWriter,
	r *http.Request,
	routeName string,
	isHx bool,
	pairs ...string,
) {
	var url string
	if strings.HasPrefix(routeName, "http") {
		url = routeName
	} else {
		url = RouteName2URL(routeName, pairs...)
	}
	if !isHx {
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		w.Header().Add("HX-Redirect", url)
	}
}
