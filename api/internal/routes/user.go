package routes

import (
	"encoding/json"
	"net/http"
	"news-api/internal"
	"news-api/internal/data"

	"golang.org/x/crypto/bcrypt"
)

type authBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, req *http.Request) {
	logger := internal.GetLogger()

	var body authBody
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Fatalln(err)
		}

		return
	}

	user, err := data.GetUserByUsername(body.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Fatalln(err)
		}

		return
	}

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		if _, err := w.Write([]byte("Invalid username or password")); err != nil {
			logger.Fatalln(err)
		}

		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if _, err := w.Write([]byte("Invalid username or password")); err != nil {
			logger.Fatalln(err)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte("Success")); err != nil {
		logger.Fatalln(err)
	}
}

func Register(w http.ResponseWriter, req *http.Request) {
	logger := internal.GetLogger()

	var body authBody
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Fatalln(err)
		}

		// Return not strictly necessary here, but if we move away from fatal logs
		return
	}

	user, err := data.GetUserByUsername(body.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Fatalln(err)
		}

		return
	} else if user != nil {
		w.WriteHeader(http.StatusConflict)
		if _, err := w.Write([]byte("User already exists")); err != nil {
			logger.Fatalln(err)
		}

		return
	}

	hashedPassword, err := hash(body.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Fatalln(err)
		}

		return
	}

	if err = data.InsertUser(body.Username, hashedPassword); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Fatalln(err)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func hash(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 16)

	return string(bytes), err
}
