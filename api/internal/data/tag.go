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
