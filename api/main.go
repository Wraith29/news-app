package main

import (
	"net/http"
	"news-api/internal/config"
	"news-api/internal/logging"
	"news-api/internal/middleware"
	"news-api/internal/routes"
)

type router struct {
	mux http.ServeMux
}

func newRouter() router {
	return router{*http.NewServeMux()}
}

func (r *router) addRoute(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(
		path,
		middleware.LoggingMiddleware(
			handler,
		),
	)
}

func (r *router) addAuthenticatedRoute(path string, handler http.HandlerFunc) {
	r.addRoute(path, middleware.AuthMiddleware(handler))
}

func main() {
	// Can't log before getting the logger config lol
	if err := config.Load(); err != nil {
		panic(err)
	}

	logger := logging.GetLogger()

	router := newRouter()
	router.addRoute("POST /auth/login", routes.Login)
	router.addRoute("POST /auth/register", routes.Register)
	router.addAuthenticatedRoute("GET /feeds", routes.GetFeeds)

	logger.Info("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", &router.mux); err != nil {
		logger.Err(err.Error())
		return
	}
}
