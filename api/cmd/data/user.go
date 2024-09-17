package data

import (
	"database/sql"

	"github.com/Wraith29/news-app/api/cmd/config"
	"github.com/Wraith29/news-app/api/cmd/models"
)

func GetUserByName(name string) (*models.User, error) {
	db, err := getConn(config.Cfg)

	if err != nil {
		return nil, err
	}

	result := db.QueryRow("SELECT username, password FROM \"user\" WHERE username = $1", name)

	var username, password string

	err = result.Scan(&username, &password)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Username: username,
		Password: password,
	}, nil
}

func CreateUser(username, password string) error {
	db, err := getConn(config.Cfg)

	if err != nil {
		return err
	}

	result := db.QueryRow("INSERT INTO \"user\" (username, password) VALUES ($1, $2)", username, password)

	err = result.Scan()

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}
