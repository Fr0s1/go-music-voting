package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	logging "user-grpc/pkg/logging"
)

var (
	Db     *sql.DB
	logger = logging.Log.WithFields(logging.StandardFields)
)

func InitDB() {
	db, err := sql.Open("mysql", "root:hieu2203@tcp(localhost)/musicvoting")

	if err != nil {
		logger.Error(err)
	}

	if err = db.Ping(); err != nil {
		logger.Panic(err)
	}

	Db = db

	logger.Info("Connected to database")
}

func CloseDB() error {
	return Db.Close()
}
