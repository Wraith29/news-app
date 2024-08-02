package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Wraith29/news-app/api/cmd/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")

		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			return
		}

		token, err := getToken(authToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		username, err := token.Claims.GetSubject()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newToken, err := createAuthToken(username)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Authorization", newToken)

		println(newToken)

		c.Next()

	}
}

// Token is nil if it is invalid (Including Expired)
func getToken(authToken string) (*jwt.Token, error) {
	claims := jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method %s. Expected %s", token.Method.Alg(), jwt.SigningMethodHS256.Name)
		}

		return []byte(config.Cfg.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	expiry, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if expiry.Before(time.Now()) {
		return nil, errors.New("AuthToken is expired")
	}

	return token, nil
}
