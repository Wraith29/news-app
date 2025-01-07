package data

import (
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type Article struct {
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	Link            string     `json:"link"`
	PublishedParsed *time.Time `json:"publishedParsed"`
	Author          string     `json:"author"`
}

func GetArticlesFromFeeds(feeds []*UserFeed) ([]*Article, error) {
	articles := make([]*Article, 0)

	parser := gofeed.NewParser()

	for _, feed := range feeds {
		if !feed.Enabled {
			continue
		}

		data, err := parser.ParseURL(feed.FeedUrl)

		if err != nil {
			return nil, err
		}

		for _, article := range data.Items {
			articles = append(articles, &Article{
				Title:           stripNewlines(article.Title),
				Description:     article.Description,
				Link:            article.Link,
				PublishedParsed: article.PublishedParsed,
				Author:          feed.FeedAuthor,
			})
		}
	}

	return articles, nil
}

func stripNewlines(s string) string {
	var left, right string
	found := true

	for found {
		left, right, found = strings.Cut(s, "\r\n")

		s = left + right
	}

	return s
}
