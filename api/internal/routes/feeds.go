package routes

import (
	"net/http"
)

func GetFeeds(w http.ResponseWriter, req *http.Request) {
	// logger := logging.GetLogger()

	// userId := req.Context().Value(ctx.ContextKeyUserId).(int)

	w.WriteHeader(http.StatusOK)
}
