package main

import (
	"net/http"
	"news-api/internal"
	"news-api/internal/config"
	"news-api/internal/middleware"
	"news-api/internal/routes"
)

func main() {
	logger := internal.GetLogger()

	if err := config.Load(); err != nil {
		logger.Fatalln(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/auth/login",
		middleware.LoggingMiddleware(
			middleware.MethodMiddleware(
				"POST", http.HandlerFunc(routes.Login),
			),
		),
	)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Fatalln(err)
	}
}
