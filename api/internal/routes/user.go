package routes

import (
	"encoding/json"
	"net/http"
	"news-api/internal/data"
	"news-api/internal/logging"

	"golang.org/x/crypto/bcrypt"
)

type authBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var body authBody
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			_ = logger.Err(err.Error())
		}

		return
	}

	user, err := data.GetUserByUsername(body.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			_ = logger.Err(err.Error())
		}

		return
	}

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		if _, err := w.Write([]byte("Invalid username or password")); err != nil {
			_ = logger.Err(err.Error())
		}

		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if _, err := w.Write([]byte("Invalid username or password")); err != nil {
			_ = logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func Register(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var body authBody
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			_ = logger.Err(err.Error())
		}

		// Return not strictly necessary here, but if we move away from fatal logs
		return
	}

	user, err := data.GetUserByUsername(body.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			_ = logger.Err(err.Error())
		}

		return
	} else if user != nil {
		w.WriteHeader(http.StatusConflict)
		if _, err := w.Write([]byte("User already exists")); err != nil {
			_ = logger.Err(err.Error())
		}

		return
	}
	hashedPassword, err := hash(body.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			_ = logger.Err(err.Error())
		}

		return
	}

	if err = data.InsertUser(body.Username, hashedPassword); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			_ = logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func hash(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 2)

	return string(bytes), err
}
