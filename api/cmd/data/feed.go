package data

import (
	"sort"

	"github.com/wraith29/news-app/api/cmd/config"
	"github.com/wraith29/news-app/api/cmd/models"
)

func GetAllFeeds(cfg *config.Config) ([]models.Feed, error) {
	db, err := getConn(cfg)

	if err != nil {
		return nil, err
	}

	feeds := make([]models.Feed, 0)

	rows, err := db.Query("SELECT id, author, feed_url FROM news_feed")

	if err != nil {
		return nil, err
	}

	var newsFeedId int
	var newsFeedAuthor string
	var newsFeedUrl string

	for rows.Next() {
		err = rows.Scan(&newsFeedId, &newsFeedAuthor, &newsFeedUrl)

		if err != nil {
			return nil, err
		}

		feeds = append(feeds, models.NewFeed(newsFeedId, newsFeedAuthor, newsFeedUrl))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return feeds, nil
}

func GetAllAuthors(cfg *config.Config) ([]string, error) {
	db, err := getConn(cfg)

	if err != nil {
		return nil, err
	}

	authors := make([]string, 0)

	rows, err := db.Query("SELECT author FROM news_feed")

	if err != nil {
		return nil, err
	}

	var author string

	for rows.Next() {
		err = rows.Scan(&author)

		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	sort.Strings(authors)

	return authors, nil
}
