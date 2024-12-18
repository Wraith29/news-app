package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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
