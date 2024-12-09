package middlewares

import "net/http"

func ForceHTMLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if h != nil {
			h.ServeHTTP(w, r)
		}
	})
}
