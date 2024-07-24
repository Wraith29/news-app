package data

import (
	"sort"

	"github.com/Wraith29/news-app/api/cmd/config"
)

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
