package main

import (
	"net/http"
	"news-api/cmd/api"
	"news-api/internal/config"
	"news-api/internal/logging"
	"news-api/internal/middleware"
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
	router.addRoute("POST /auth/login", api.Login)
	router.addRoute("POST /auth/register", api.Register)
	router.addAuthenticatedRoute("POST /feed", api.AddFeed)
	router.addAuthenticatedRoute("GET /feed", api.GetUserFeeds)
	router.addAuthenticatedRoute("PUT /feed/join", api.JoinFeed)
	router.addAuthenticatedRoute("PUT /feed/leave", api.LeaveFeed)
	router.addAuthenticatedRoute("POST /tag", api.AddTag)
	router.addAuthenticatedRoute("DELETE /tag", api.DeleteTag)
	router.addAuthenticatedRoute("GET /feed/tag", api.GetTagsForFeed)
	router.addAuthenticatedRoute("PUT /feed/tag", api.TagFeed)
	router.addAuthenticatedRoute("DELETE /feed/tag", api.UnTagFeed)
	router.addAuthenticatedRoute("GET /articles", api.GetArticles)

	logger.Info("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", &router.mux); err != nil {
		logger.Err(err.Error())
		return
	}
}
