package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var _db *sql.DB

func getConn() (*sql.DB, error) {
	if _db != nil {
		return _db, nil
	}

	connStr := "user=postgres password=password dbname=news_feed sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		_db = nil
		fmt.Printf("Failed to open DB: %s\n", err.Error())
		return nil, err
	}

	_db = db
	return _db, nil
}
