package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	invalidToken = errors.New("invalid token signing")
	expiredToken = errors.New("expired token")
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte("Missing required header \"Authorization\"")); err != nil {
				panic(err)
			}

			return
		}

		token, err := getAuthToken(authHeader)
		if err != nil && err == expiredToken || err == invalidToken {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				panic(err)
			}

			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				panic(err)
			}

			return
		}

		usn, err := token.Claims.GetSubject()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil {
				panic(err)
			}

			return
		}

		println(usn)

		next.ServeHTTP(w, req)
	})
}

func getAuthToken(tkn string) (*jwt.Token, error) {
	claims := jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tkn, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidToken
		}

		return []byte("secret-key"), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	expiry, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if expiry.Before(time.Now()) {
		return nil, expiredToken
	}

	return token, nil
}
