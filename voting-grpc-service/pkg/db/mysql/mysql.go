package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	logging "voting-grpc/pkg/logging"
)

var Db *sql.DB

func InitDB() {
	logger := logging.Log.WithFields(logging.StandardFields)
	// logger := logging.Log

	db, err := sql.Open("mysql", "root:hieu2203@tcp(localhost)/musicvoting")

	if err != nil {
		logger.Error(err)
	}

	if err = db.Ping(); err != nil {
		logger.Error(err)
	}

	Db = db

	logger.Info("Connected to database")
}

func CloseDB() error {
	return Db.Close()
}
