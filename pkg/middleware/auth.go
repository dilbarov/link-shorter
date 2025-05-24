package middleware

import (
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer = strings.TrimPrefix(bearer, "Bearer ")

		next.ServeHTTP(w, r)
	})
}
