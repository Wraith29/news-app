package api

import (
	"encoding/json"
	"net/http"
	"news-api/internal/ctx"
	"news-api/internal/data"
	"news-api/internal/logging"
)

func AddTag(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var body struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	if err := data.AddTag(body.Name); err != nil {
		data.HandleDataError(w, req, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTag(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var body struct {
		Id int `json:"id"`
	}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	if err := data.DeleteTag(body.Id); err != nil {
		data.HandleDataError(w, req, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func TagFeed(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	var body struct {
		FeedId int `json:"feedId"`
		TagId  int `json:"tagId"`
	}

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	success, err := data.TagFeed(body.FeedId, body.TagId, userId)
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

	w.WriteHeader(http.StatusNoContent)
}
