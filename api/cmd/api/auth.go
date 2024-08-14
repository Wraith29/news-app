package api

import (
	"net/http"
	"time"

	"github.com/Wraith29/news-app/api/cmd/config"
	"github.com/Wraith29/news-app/api/cmd/data"
	"github.com/Wraith29/news-app/api/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func login(c *gin.Context) {
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

	token, err := createAuthToken(user.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authToken": token})
	return

}

func register(c *gin.Context) {
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

func createAuthToken(username string) (string, error) {
	now := time.Now()

	claims := jwt.RegisteredClaims{
		Issuer:    "news-feed-app",
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(now),
		Audience:  jwt.ClaimStrings{"http://localhost:2912"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.Cfg.SecretKey))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
