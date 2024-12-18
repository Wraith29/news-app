package data

type Feed struct {
	Id         int    `json:"id"`
	FeedAuthor string `json:"feed_author"`
	FeedUrl    string `json:"feed_url"`
	AddedBy    int    `json:"added_by"`
}

func AddFeed(author, url string, addedBy int) error {
	conn, err := getDb()
	if err != nil {
		return err
	}

	query := `INSERT INTO "news_feed" ("feed_author", "feed_url", "added_by") VALUES ($1, $2, $3)`

	_, err = conn.Exec(query, author, url, addedBy)

	return err
}

func GetUserFeeds(userId int) ([]*Feed, error) {
	conn, err := getDb()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT NF."id", "feed_author", "feed_url", "added_by"
		FROM "news_feed" NF
		INNER JOIN "user_feed" UF ON NF."id" = UF."feed_id"
		WHERE UF."user_id" = $1
	`

	feeds, err := conn.Query(query, userId)
	if err != nil {
		return nil, err
	}

	result := make([]*Feed, 0)

	for feeds.Next() {
		feed := Feed{}

		if err := feeds.Scan(&feed.Id, &feed.FeedAuthor, &feed.FeedUrl, &feed.AddedBy); err != nil {
			return nil, err
		}

		result = append(result, &feed)
	}

	if err := feeds.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func JoinFeed(feedId, userId int) (bool, error) {
	conn, err := getDb()
	if err != nil {
		return false, err
	}

	query := `
		INSERT INTO "user_feed" ("user_id", "feed_id")
		VALUES ($1, $2)
	`

	result, err := conn.Exec(query, userId, feedId)

	if err != nil {
		return false, err
	}

	rowCount, err := result.RowsAffected()

	// If rowCount == 1 then the user has joined the feed, otherwise there's a problem
	return rowCount == 1, err
}

func LeaveFeed(feedId, userId int) (bool, error) {
	conn, err := getDb()
	if err != nil {
		return false, err
	}

	query := `
		DELETE FROM "user_feed"
		WHERE "user_id" = $1 AND "feed_id" = $2
	`

	result, err := conn.Exec(query, userId, feedId)

	if err != nil {
		return false, err
	}

	rowCount, err := result.RowsAffected()
	return rowCount == 1, err
}
