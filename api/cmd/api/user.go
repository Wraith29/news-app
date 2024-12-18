package api

import (
	"encoding/json"
	"net/http"
	"news-api/internal/config"
	"news-api/internal/data"
	"news-api/internal/logging"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	AuthToken string `json:"authToken"`
}

func Login(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var body authPayload

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	user, err := data.GetUserByUsername(body.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		if _, err := w.Write([]byte("Invalid username or password")); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if _, err := w.Write([]byte("Invalid username or password")); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	token, err := getAuthToken(user.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	authRes := authResponse{token}

	response, err := json.Marshal(authRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logger.Err(err.Error())
	}
}

func Register(w http.ResponseWriter, req *http.Request) {
	logger := logging.GetLogger()

	var body authPayload
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	user, err := data.GetUserByUsername(body.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	} else if user != nil {
		w.WriteHeader(http.StatusConflict)
		if _, err := w.Write([]byte("User already exists")); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	hashedPassword, err := hash(body.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	userId, err := data.InsertUser(body.Username, hashedPassword)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	token, err := getAuthToken(userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	authRes := authResponse{token}

	response, err := json.Marshal(authRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logger.Err(err.Error())
	}
}

func hash(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 2)

	return string(bytes), err
}

func getAuthToken(userId int) (string, error) {
	now := time.Now()

	claims := jwt.RegisteredClaims{
		Issuer:    "news-feed-api",
		Subject:   strconv.Itoa(userId),
		ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(now),
		Audience:  jwt.ClaimStrings{"http://localhost:2912"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Cfg.SecretKey))
}
