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

func (r *router) addRoute(path, method string, handler http.HandlerFunc) {
	r.mux.HandleFunc(
		path,
		middleware.LoggingMiddleware(
			middleware.MethodMiddleware(
				method, handler,
			),
		),
	)
}

func (r *router) addAuthenticatedRoute(path, method string, handler http.HandlerFunc) {
	r.addRoute(path, method, middleware.AuthMiddleware(handler))
}

func main() {
	// Can't log before getting the logger config lol
	if err := config.Load(); err != nil {
		panic(err)
	}

	logger := logging.GetLogger()

	router := newRouter()
	router.addRoute("/auth/login", "POST", routes.Login)
	router.addRoute("/auth/register", "POST", routes.Register)

	_ = logger.Info("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", &router.mux); err != nil {
		_ = logger.Err(err.Error())
		return
	}
}
