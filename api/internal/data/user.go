package data

import "database/sql"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserByUsername(username string) (*User, error) {
	conn, err := GetDb()

	if err != nil {
		return nil, err
	}

	query := `SELECT "id", "username", "password" FROM "user" WHERE "username" = $1`

	result := conn.QueryRow(query, username)

	user := User{}

	err = result.Scan(&user.Id, &user.Username, &user.Password)

	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func InsertUser(username, password string) error {
	conn, err := GetDb()

	if err != nil {
		return err
	}

	query := `INSERT INTO "user" ("username", "password") VALUES ($1, $2)`

	_, err = conn.Exec(query, username, password)

	return err
}
