package middleware

import (
	"context"
	"link-shorter/configs"
	"link-shorter/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextUserKey key = "ContextUserKey"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	_, err := w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
	if err != nil {
		return
	}
}

func IsAuthed(next http.Handler, conf *configs.AuthConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			writeUnauthorized(w)
			return
		}

		if !strings.HasPrefix(bearer, "Bearer ") {
			writeUnauthorized(w)
			return
		}

		bearer = strings.TrimPrefix(bearer, "Bearer ")

		isValid, data := jwt.NewJWTService(conf.Secret).Parse(bearer)

		if !isValid {
			writeUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, data)

		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
