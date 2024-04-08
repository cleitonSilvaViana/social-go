package middleware

import (
	"log/slog"
	"net/http"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		slog.LogAttrs(
			r.Context(),
			slog.LevelInfo, "Start request",
			slog.String("URL", r.URL.String()),
			slog.String("method", r.Method),
		)

		next.ServeHTTP(w, r)

		slog.LogAttrs(
			r.Context(),
			slog.LevelInfo, "Completed request",
			slog.String("URL", r.URL.String()),
			// slog.Int("status_code", 200),
		)

	})
}
