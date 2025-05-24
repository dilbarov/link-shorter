package middleware

import (
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func Logging(next http.Handler, logger zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{ResponseWriter: w, StatusCode: http.StatusOK}

		next.ServeHTTP(wrapper, r)

		logger.Info().Msgf("%d %s %s | %dms", wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start).Milliseconds())
	})
}

func NewLoggingMiddleware(logger zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return Logging(next, logger)
	}
}
