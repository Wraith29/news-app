package middleware

import (
	"net/http"
	"news-api/internal"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	logger := internal.GetLogger()

	return func(w http.ResponseWriter, req *http.Request) {
		logger.Printf("%s %s\n", req.URL.Path, req.Method)

		next.ServeHTTP(w, req)
	}
}
