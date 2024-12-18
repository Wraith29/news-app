package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news-api/internal/ctx"
	"news-api/internal/data"
	"news-api/internal/logging"
)

type addTagRequest struct {
	Name string `json:"name"`
}

func AddTag(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var tagRequest addTagRequest

	logger.Info("Decoding body into Add Tag request")
	if err := json.NewDecoder(req.Body).Decode(&tagRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info(fmt.Sprintf("Creating new tag %s", tagRequest.Name))

	if err := data.AddTag(tagRequest.Name); err != nil {
		data.HandleDataError(w, req, err)
		return
	}

	logger.Info(fmt.Sprintf("Tag %s created", tagRequest.Name))

	w.WriteHeader(http.StatusNoContent)
}

type deleteTagRequest struct {
	Id int `json:"id"`
}

func DeleteTag(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var deleteRequest deleteTagRequest
	logger.Info("Decoding body into Delete Tag request")

	if err := json.NewDecoder(req.Body).Decode(&deleteRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info(fmt.Sprintf("Deleting tag %d", deleteRequest.Id))

	if err := data.DeleteTag(deleteRequest.Id); err != nil {
		data.HandleDataError(w, req, err)
		return
	}

	logger.Info(fmt.Sprintf("Tag %d successfully deleted", deleteRequest.Id))

	w.WriteHeader(http.StatusNoContent)
}

type tagFeedRequest struct {
	FeedId int `json:"feedId"`
	TagId  int `json:"tagId"`
}

func TagFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	var tagRequest tagFeedRequest
	logger.Info("Decoding body into tagFeedRequest")

	if err := json.NewDecoder(req.Body).Decode(&tagRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info(fmt.Sprintf("User %d adding tag %d to feed %d", userId, tagRequest.TagId, tagRequest.FeedId))

	success, err := data.TagFeed(tagRequest.FeedId, tagRequest.TagId, userId)
	if err != nil {
		data.HandleDataError(w, req, err)
		return
	}

	if !success {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("failed to tag feed")); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	logger.Info(fmt.Sprintf("Tagged feed %d with %d", tagRequest.FeedId, tagRequest.TagId))

	w.WriteHeader(http.StatusNoContent)
}
