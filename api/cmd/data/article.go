package data

import (
	"sort"

	"github.com/mmcdole/gofeed"
	"github.com/wraith29/news-app/api/cmd/models"
)

func GetAllArticles() (models.ArticleList, error) {
	feeds, err := GetAllFeeds()

	if err != nil {
		return nil, err
	}

	parser := gofeed.NewParser()
	articles := make(models.ArticleList, 0)

	for _, feed := range feeds {
		newsFeed, err := parser.ParseURL(feed.FeedUrl)

		if err != nil {
			return nil, err
		}

		for _, article := range newsFeed.Items {
			articles = append(articles, models.Article{
				Title:           article.Title,
				Description:     article.Description,
				Link:            article.Link,
				PublishedParsed: article.PublishedParsed,
				Author:          feed.Author,
			})
		}
	}

	sort.Sort(articles)

	return articles, nil
}
