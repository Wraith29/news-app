package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Wraith29/news-app/api/cmd/config"
	"github.com/Wraith29/news-app/api/cmd/data"
	"github.com/Wraith29/news-app/api/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func authenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")

		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			return
		}

		claims := jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				println("Method is not HMAC")
				return nil, fmt.Errorf("Invalid signing method %s. Expected %s", token.Method.Alg(), jwt.SigningMethodHS256.Name)
			}

			return []byte(config.Cfg.SecretKey), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("Invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		expiry, err := claims.GetExpirationTime()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if expiry.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
			return
		}

		c.Next()
	}
}

func authUser(c *gin.Context) {
	incoming := models.User{}

	err := c.ShouldBindJSON(&incoming)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := data.GetUserByName(incoming.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or Password"})
		return
	}

	if user.Password != incoming.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or Password"})
		return
	}

	now := time.Now()

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(now),
		Audience:  jwt.ClaimStrings{"http://localhost:2912"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.Cfg.SecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authToken": tokenStr})
	return

}

func createUser(c *gin.Context) {
	incoming := models.User{}

	err := c.ShouldBindJSON(&incoming)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := data.CreateUser(incoming.Username, incoming.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, id)
}
