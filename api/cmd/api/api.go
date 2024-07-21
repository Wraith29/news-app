package api

import (
	"crypto/sha256"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wraith29/news-app/api/cmd/data"
	"github.com/wraith29/news-app/api/cmd/models"
)

type cacheableRequest[T any] struct {
	Value T      `json:"value"`
	Hash  []byte `json:"hash"`
}

func AddApiRoutes(e *gin.Engine) {
	e.GET("/feeds", getAllFeeds)
	e.GET("/articles", getAllArticles)
	e.GET("/authors", getAllAuthors)
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

func getAllAuthors(c *gin.Context) {
	authors, err := data.GetAllAuthors()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	authorHash := sha256.New()
	hash := make([]byte, 0)

	for _, author := range authors {
		authorHash.Write([]byte(author))

		for _, b := range authorHash.Sum(nil) {
			hash = append(hash, b)
		}
	}

	request := cacheableRequest[[]string]{
		Value: authors,
		Hash:  hash,
	}

	c.JSON(http.StatusOK, request)
}
