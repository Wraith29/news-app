package api

import (
	"net/http"
	"rss-app/cmd/data"
	"rss-app/cmd/models"

	"github.com/gin-gonic/gin"
)

type cacheableRequest[T any] struct {
	Value T      `json:"value"`
	Hash  []byte `json:"hash"`
}

func AddApiRoutes(e *gin.Engine) {
	e.GET("/feeds", getAllFeeds)
	e.GET("/articles", getAllArticles)
}

func getAllFeeds(c *gin.Context) {
	feeds, err := data.GetAllFeeds()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feeds)
}

func getAllArticles(c *gin.Context) {
	articles, err := data.GetAllArticles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request := cacheableRequest[[]models.Article]{
		Value: articles,
		Hash:  articles.Hash(),
	}

	c.JSON(http.StatusOK, request)
}
