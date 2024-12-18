package api

import (
	"encoding/json"
	"fmt"
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

	logger.Info("Decoding body into Feed Request")
	if err := json.NewDecoder(req.Body).Decode(&feedRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info(fmt.Sprintf("Adding New Feed %s", feedRequest.FeedAuthor))
	if err := data.AddFeed(feedRequest.FeedAuthor, feedRequest.FeedUrl, userId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info("Feed Added Succesfully")
	w.WriteHeader(http.StatusNoContent)
}

func GetUserFeeds(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	logger.Info(fmt.Sprintf("Retrieving feeds for user %d", userId))
	feeds, err := data.GetUserFeeds(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info("Loading feeds into response")
	msg, err := json.Marshal(&feeds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info("Feeds Loaded Successfully")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(msg); err != nil {
		logger.Err(err.Error())
	}
}

type feedIdRequest struct {
	FeedId int `json:"feedId"`
}

func JoinFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	logger.Info(fmt.Sprintf("User %d attempting to join feed", userId))

	var feedRequest feedIdRequest
	if err := json.NewDecoder(req.Body).Decode(&feedRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info(fmt.Sprintf("User %d joining %d", userId, feedRequest.FeedId))
	success, err := data.JoinFeed(feedRequest.FeedId, userId)

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

	logger.Info(fmt.Sprintf("User %d successfully joined %d", userId, feedRequest.FeedId))

	w.WriteHeader(http.StatusNoContent)
}

func LeaveFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	logger.Info(fmt.Sprintf("User %d attempting to leave feed", userId))

	var feedRequest feedIdRequest
	if err := json.NewDecoder(req.Body).Decode(&feedRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info(fmt.Sprintf("User %d leaving feed %d", userId, feedRequest.FeedId))
	success, err := data.LeaveFeed(feedRequest.FeedId, userId)

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

	logger.Info(fmt.Sprintf("User %d successfully joined %d", userId, feedRequest.FeedId))

	w.WriteHeader(http.StatusNoContent)
}
