package middleware

import (
	"fmt"
	"net/http"
	"news-api/internal/logging"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	logger := logging.GetLogger()

	return func(w http.ResponseWriter, req *http.Request) {
		logger.Info(fmt.Sprintf("%s %s", req.URL.Path, req.Method))

		next.ServeHTTP(w, req)
	}
}
