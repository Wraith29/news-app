package main

import (
	"net/http"
	"news-api/internal/config"
	"news-api/internal/logging"
	"news-api/internal/middleware"
	"news-api/internal/routes"
)

func main() {
	logger := logging.GetLogger()

	if err := config.Load(); err != nil {
		logger.Err(err.Error())
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/auth/login",
		middleware.LoggingMiddleware(
			middleware.MethodMiddleware(
				"POST",
				http.HandlerFunc(routes.Login),
			),
		),
	)

	mux.HandleFunc(
		"/auth/register",
		middleware.LoggingMiddleware(
			middleware.MethodMiddleware(
				"POST",
				http.HandlerFunc(routes.Register),
			),
		),
	)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Err(err.Error())
		return
	}
}
