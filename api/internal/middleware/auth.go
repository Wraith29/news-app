package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"news-api/internal/config"
	"news-api/internal/ctx"
	"news-api/internal/logging"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	invalidToken = errors.New("invalid token signing")
	expiredToken = errors.New("expired token")
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger := logging.GetLogger()
		authHeader := req.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte("Missing required header \"Authorization\"")); err != nil {
				logger.Err(err.Error())
			}

			return
		}

		token, err := getAuthToken(authHeader)
		if err != nil && err == expiredToken || err == invalidToken {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				logger.Err(err.Error())
			}

			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				logger.Err(err.Error())
			}

			return
		}

		rawId, err := token.Claims.GetSubject()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				logger.Err(err.Error())
			}

			return
		}

		userId, err := strconv.Atoi(rawId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				logger.Err(err.Error())
			}

			return
		}

		userContext := context.WithValue(req.Context(), ctx.ContextKeyUserId, userId)

		next.ServeHTTP(w, req.WithContext(userContext))
	})
}

func getAuthToken(tkn string) (*jwt.Token, error) {
	logger := logging.GetLogger()
	logger.Info(fmt.Sprintf("Received token %s", tkn))

	claims := jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tkn, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidToken
		}

		return []byte(config.Cfg.SecretKey), nil
	})

	logger.Info("Parsed claims from token")

	if err != nil || !token.Valid {
		return nil, err
	}

	logger.Info("Getting claims expiration date")

	expiry, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if expiry.Before(time.Now()) {
		return nil, expiredToken
	}

	return token, nil
}
