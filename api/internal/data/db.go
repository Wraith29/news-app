package data

import (
	"database/sql"
	"fmt"
	"news-api/internal"
	"news-api/internal/config"

	_ "github.com/lib/pq"
)

var _db *sql.DB = nil

func GetDb() (*sql.DB, error) {
	logger := internal.GetLogger()

	if _db != nil {
		return _db, nil
	}

	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=news_feed sslmode=disable host=%s port=%d",
		config.Cfg.Postgres.Username, config.Cfg.Postgres.Password,
		config.Cfg.Postgres.Host, config.Cfg.Postgres.Port,
	)

	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		logger.Println(err)
		_db = nil
		return nil, err
	}

	_db = conn
	return _db, nil
}
