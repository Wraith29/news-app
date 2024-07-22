package api

import (
	"crypto/sha256"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wraith29/news-app/api/cmd/config"
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
	feeds, err := data.GetAllFeeds(config.Cfg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	feedAuthors := make([]string, 0)
	for _, feed := range feeds {
		feedAuthors = append(feedAuthors, feed.Author)
	}

	request := cacheableRequest[[]models.Feed]{
		Value: feeds,
		Hash:  hashStrings(feedAuthors),
	}

	c.JSON(http.StatusOK, request)
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
	authors, err := data.GetAllAuthors(config.Cfg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hash := hashStrings(authors)

	request := cacheableRequest[[]string]{
		Value: authors,
		Hash:  hash,
	}

	c.JSON(http.StatusOK, request)
}

func hashStrings(inputs []string) []byte {
	hash := sha256.New()
	sum := make([]byte, 0)

	for _, input := range inputs {
		hash.Write([]byte(input))

		for _, b := range hash.Sum(nil) {
			sum = append(sum, b)
		}
	}

	return sum
}
