package data

type Feed struct {
	Id          int    `json:"id"`
	FeedAuthor  string `json:"feedAuthor"`
	FeedUrl     string `json:"feedUrl"`
	AddedBy     int    `json:"addedBy"`
	Subscribers int    `json:"subscribers"`
	IsJoined    bool   `json:"isJoined"`
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

func GetAllFeeds(userId int) ([]*Feed, error) {
	conn, err := getDb()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT NF."id",
				NF."feed_author",
				NF."feed_url",
				NF."added_by",
				(SELECT CAST(COUNT(*) AS INT) AS "subscribers" FROM "user_feed" WHERE "feed_id" = NF."id"),
				(SELECT EXISTS(SELECT 1 FROM "user_feed" WHERE "feed_id" = NF."id" AND "user_id" = $1))
		FROM "news_feed" NF
	`

	feeds, err := conn.Query(query, userId)
	if err != nil {
		return nil, err
	}

	result := make([]*Feed, 0)

	for feeds.Next() {
		feed := Feed{}

		if err := feeds.Scan(&feed.Id, &feed.FeedAuthor, &feed.FeedUrl, &feed.AddedBy, &feed.Subscribers, &feed.IsJoined); err != nil {
			return nil, err
		}

		result = append(result, &feed)
	}

	return result, feeds.Err()
}

type UserFeed struct {
	Feed
	Enabled bool `json:"enabled"`
}

func GetUserFeeds(userId int) ([]*UserFeed, error) {
	conn, err := getDb()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT NF."id", NF."feed_author",
				NF."feed_url", NF."added_by",
				UF."enabled", (SELECT CAST(COUNT(*) AS INT) AS "subscribers" FROM "user_feed" WHERE "feed_id" = NF."id")
		FROM "news_feed" NF
		INNER JOIN "user_feed" UF
		ON NF."id" = UF."feed_id"
		WHERE UF."user_id" = $1;
	`

	feeds, err := conn.Query(query, userId)
	if err != nil {
		return nil, err
	}

	result := make([]*UserFeed, 0)

	for feeds.Next() {
		feed := UserFeed{}

		if err := feeds.Scan(&feed.Id, &feed.FeedAuthor, &feed.FeedUrl, &feed.AddedBy, &feed.Enabled, &feed.Subscribers); err != nil {
			return nil, err
		}

		result = append(result, &feed)
	}

	return result, feeds.Err()
}

func ToggleFeed(feedId, userId int) (bool, error) {
	conn, err := getDb()
	if err != nil {
		return false, err
	}

	query := `
		UPDATE "user_feed"
		SET "enabled" = NOT "enabled"
		WHERE "user_id" = $1 AND "feed_id" = $2
	`

	result, err := conn.Exec(query, userId, feedId)

	if err != nil {
		return false, err
	}

	rowCount, err := result.RowsAffected()

	return rowCount == 1, err
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

func GetTagsForFeed(feedId int) ([]Tag, error) {
	conn, err := getDb()
	if err != nil {
		return nil, err
	}

	query := `
		SELECT T."id", T."name"
		FROM "tag" T
		INNER JOIN "feed_tag" FT
			ON FT."feed_id" = $1
	`

	results, err := conn.Query(query, feedId)
	if err != nil {
		return nil, err
	}

	tags := make([]Tag, 0)

	for results.Next() {
		var tag Tag

		if err := results.Scan(&tag.Id, &tag.Name); err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, results.Err()
}
