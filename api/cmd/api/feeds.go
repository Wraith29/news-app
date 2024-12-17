package api

import (
	"encoding/json"
	"net/http"
	"news-api/internal/ctx"
	"news-api/internal/data"
	"news-api/internal/logging"
)

type addFeedRequest struct {
	FeedAuthor string `json:"feed_author"`
	FeedUrl    string `json:"feed_url"`
}

func AddFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	var feedRequest addFeedRequest

	if err := json.NewDecoder(req.Body).Decode(&feedRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	if err := data.AddFeed(feedRequest.FeedAuthor, feedRequest.FeedUrl, userId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetUserFeeds(w http.ResponseWriter, req *http.Request) {
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

	msg, err := json.Marshal(&feeds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(msg); err != nil {
		logger.Err(err.Error())
	}
}
