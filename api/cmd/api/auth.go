package api

import (
	"net/http"

	"github.com/Wraith29/news-app/api/cmd/data"
	"github.com/Wraith29/news-app/api/cmd/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

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

	token := jwt.New(jwt.SigningMethodHS256)

	tokenStr, err := token.SignedString("my-secret")

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
