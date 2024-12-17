package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MethodMiddleware(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			req.Body.Close()

			w.WriteHeader(http.StatusBadRequest)

			res, _ := json.Marshal(struct {
				Error string `json:"error"`
			}{
				Error: fmt.Sprintf("Invalid Method. Expected %s got %s.", method, req.Method),
			})

			if _, err := w.Write(res); err != nil {
				panic(err)
			}
			return
		}

		next.ServeHTTP(w, req)
	}
}
