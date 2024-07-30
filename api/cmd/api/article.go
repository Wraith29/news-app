package api

import (
	"net/http"

	"github.com/Wraith29/news-app/api/cmd/data"
	"github.com/Wraith29/news-app/api/cmd/models"
	"github.com/gin-gonic/gin"
)

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
