package data

import (
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

func GetArticlesFromFeeds(feeds []*Feed) ([]*Article, error) {
	articles := make([]*Article, 0)

	parser := gofeed.NewParser()

	for _, feed := range feeds {
		data, err := parser.ParseURL(feed.FeedUrl)

		if err != nil {
			return nil, err
		}

		for _, article := range data.Items {
			articles = append(articles, &Article{
				Title:           article.Title,
				Description:     article.Description,
				Link:            article.Link,
				PublishedParsed: article.PublishedParsed,
				Author:          feed.FeedAuthor,
			})
		}
	}

	return articles, nil
}
