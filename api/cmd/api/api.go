package api

import (
	"crypto/sha256"
	"net/http"

	"github.com/Wraith29/news-app/api/cmd/config"
	"github.com/Wraith29/news-app/api/cmd/data"
	"github.com/gin-gonic/gin"
)

type cacheableRequest[T any] struct {
	Value T      `json:"value"`
	Hash  []byte `json:"hash"`
}

func AddApiRoutes(e *gin.Engine) {
	e.POST("/login", authUser)

	authenticatedRoutes := e.Group("/")

	authenticatedRoutes.Use(authMiddleware())

	authenticatedRoutes.POST("/user", createUser)
	authenticatedRoutes.GET("/feeds", getAllFeeds)
	authenticatedRoutes.PUT("/feed", updateFeed)
	authenticatedRoutes.DELETE("/feed", deleteFeed)
	authenticatedRoutes.POST("/feed", createFeed)
	authenticatedRoutes.GET("/articles", getAllArticles)
	authenticatedRoutes.GET("/authors", getAllAuthors)
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
