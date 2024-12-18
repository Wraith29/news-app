package api

import (
	"encoding/json"
	"net/http"
	"news-api/internal/ctx"
	"news-api/internal/data"
	"news-api/internal/logging"
)

func AddFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	var body struct {
		FeedAuthor string `json:"feed_author"`
		FeedUrl    string `json:"feed_url"`
	}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	if err := data.AddFeed(body.FeedAuthor, body.FeedUrl, userId); err != nil {
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

func JoinFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	var body struct {
		FeedId int `json:"feedId"`
	}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	success, err := data.JoinFeed(body.FeedId, userId)

	if err != nil {
		data.HandleDataError(w, req, err)
		return
	}

	if !success {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("failed to join news feed")); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func LeaveFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	var body struct {
		FeedId int `json:"feedId"`
	}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	success, err := data.LeaveFeed(body.FeedId, userId)

	if err != nil {
		data.HandleDataError(w, req, err)
		return
	}

	if !success {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("failed to leave news feed")); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetTagsForFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var body struct {
		FeedId int `json:"feedId"`
	}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	tags, err := data.GetTagsForFeed(body.FeedId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	msg, err := json.Marshal(tags)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(msg)); err != nil {
		logger.Err(err.Error())
	}
}
