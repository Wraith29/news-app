package data

import (
	"database/sql"
	"fmt"
	"net/http"
	"news-api/internal/config"
	"news-api/internal/logging"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var _db *sql.DB = nil

func getDb() (*sql.DB, error) {
	logger := logging.GetLogger()

	if _db != nil {
		return _db, nil
	}

	logger.Info("SQL Connection not found. Creating one")

	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=news_feed sslmode=disable host=%s port=%d",
		config.Cfg.Postgres.Username, config.Cfg.Postgres.Password,
		config.Cfg.Postgres.Host, config.Cfg.Postgres.Port,
	)

	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		logger.Err(err.Error())
		_db = nil
		return nil, err
	}

	_db = conn
	return _db, nil
}

// Integrity constraint violation
//
// - Foreign Key violation (Invalid feed or user)
// - Unique Key violation (User already on Feed)
const integrityError = "23"

func HandleDataError(w http.ResponseWriter, req *http.Request, err error) {
	logger := logging.GetLogger()

	if err, isPqErr := err.(*pq.Error); isPqErr && err.Code.Class() == integrityError {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Code.Name())); err != nil {
			logger.Err(err.Error())
		}

		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte(err.Error())); err != nil {
		logger.Err(err.Error())
	}
}
