package data

import (
	"database/sql"
	"fmt"

	"github.com/Wraith29/news-app/api/cmd/config"
	_ "github.com/lib/pq"
)

var _db *sql.DB

func getConn(cfg *config.Config) (*sql.DB, error) {
	if _db != nil {
		return _db, nil
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=news_feed sslmode=disable host=%s",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host,
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		_db = nil
		fmt.Printf("Failed to open DB: %s\n", err.Error())
		return nil, err
	}

	_db = db
	return _db, nil
}
