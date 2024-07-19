package data

import (
	"fmt"

	"github.com/wraith29/news-app/api/cmd/models"
)

func GetAllFeeds() ([]models.Feed, error) {
	db, err := getConn()

	if err != nil {
		fmt.Printf("Failed to get conn: %s\n", err.Error())
		return nil, err
	}

	feeds := make([]models.Feed, 0)

	rows, err := db.Query("SELECT id, author, feed_url FROM news_feed")

	if err != nil {
		fmt.Printf("Failed to query DB: %s\n", err.Error())
		return nil, err
	}

	var newsFeedId int
	var newsFeedAuthor string
	var newsFeedUrl string

	for rows.Next() {
		err = rows.Scan(&newsFeedId, &newsFeedAuthor, &newsFeedUrl)

		if err != nil {
			fmt.Printf("Failed to Scan row: %s\n", err.Error())
			return nil, err
		}

		feeds = append(feeds, models.NewFeed(newsFeedId, newsFeedAuthor, newsFeedUrl))
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error scanning rows: %s\n", err.Error())
		return nil, err
	}

	return feeds, nil
}
