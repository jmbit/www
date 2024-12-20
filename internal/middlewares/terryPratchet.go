package middlewares

import (
	"net/http"
)

func TerryPratchet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Clacks-Overhead", "Terry Pratchett")
		next.ServeHTTP(w, r)
	})
}
