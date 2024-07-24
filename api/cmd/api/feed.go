package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Wraith29/news-app/api/cmd/config"
	"github.com/Wraith29/news-app/api/cmd/data"
	"github.com/Wraith29/news-app/api/cmd/models"
	"github.com/gin-gonic/gin"
)

func getAllFeeds(c *gin.Context) {
	feeds, err := data.GetAllFeeds(config.Cfg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	feedStrings := make([]string, 0)
	for _, feed := range feeds {
		feedStrings = append(feedStrings, feed.String())
	}

	request := cacheableRequest[[]models.Feed]{
		Value: feeds,
		Hash:  hashStrings(feedStrings),
	}

	c.JSON(http.StatusOK, request)
}

func updateFeed(c *gin.Context) {
	feed := models.Feed{}

	err := c.ShouldBindJSON(&feed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = data.UpdateFeed(config.Cfg, feed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func deleteFeed(c *gin.Context) {
	id, exists := c.GetQuery("feedId")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter feedId"})
		return
	}

	feedId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid FeedId %s", id)})
		return
	}

	err = data.DeleteFeed(config.Cfg, feedId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
