package data

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func AddTag(name string) error {
	conn, err := getDb()
	if err != nil {
		return err
	}

	query := `INSERT INTO "tag" ("name") VALUES ($1)`

	_, err = conn.Exec(query, name)

	return err
}

func DeleteTag(tagId int) error {
	conn, err := getDb()
	if err != nil {
		return err
	}

	query := `
		DELETE FROM "tag"
		WHERE "id" = $1
	`

	_, err = conn.Exec(query, tagId)

	return err
}

func TagFeed(feedId, tagId, userId int) (bool, error) {
	conn, err := getDb()
	if err != nil {
		return false, err
	}

	query := `
		INSERT INTO "feed_tag" ("feed_id", "tag_id", "tagged_by")
		VALUES ($1, $2, $3)
	`

	result, err := conn.Exec(query, feedId, tagId, userId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()

	return rowsAffected == 1, err
}
