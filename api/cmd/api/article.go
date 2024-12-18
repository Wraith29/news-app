package api

import (
	"encoding/json"
	"net/http"
	"news-api/internal/ctx"
	"news-api/internal/data"
	"news-api/internal/logging"
	"slices"
)

func GetArticles(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	feeds, err := data.GetUserFeeds(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	articles, err := data.GetArticlesFromFeeds(feeds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	slices.SortFunc(articles, func(l, r *data.Article) int {
		return l.PublishedParsed.UTC().Compare(r.PublishedParsed.UTC())
	})

	response, err := json.Marshal(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logger.Err(err.Error())
	}
}
