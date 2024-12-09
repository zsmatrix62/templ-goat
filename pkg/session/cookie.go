package session

import (
	"net/http"
	"time"
)

func ClearAllCookies(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		http.SetCookie(w, &http.Cookie{
			Name:     cookie.Name,
			Value:    "",
			Path:     "/",
			Domain:   "localhost",
			Expires:  time.Unix(0, 0), // 设置过期时间
			MaxAge:   -1,
			HttpOnly: true,
		})
	}
}
