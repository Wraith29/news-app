package data

import (
	"github.com/Wraith29/news-app/api/cmd/config"
	"github.com/Wraith29/news-app/api/cmd/models"
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

func UpdateFeed(cfg *config.Config, feed models.Feed) error {
	db, err := getConn(cfg)

	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE news_feed SET author = $1, feed_url = $2 WHERE id = $3", feed.Author, feed.FeedUrl, feed.Id)

	return err
}

func DeleteFeed(cfg *config.Config, id int) error {
	db, err := getConn(cfg)

	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM news_feed WHERE id = $1", id)

	return err
}
