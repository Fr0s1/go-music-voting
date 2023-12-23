package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	logging "music-service/pkg/logging"
)

var Db *sql.DB

func InitDB() {
	log := logging.Log.WithFields(logging.StandardFields)

	db, err := sql.Open("mysql", "root:hieu2203@tcp(localhost)/musicvoting")

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	Db = db

	log.Info("Connected")
}

func CloseDB() error {
	return Db.Close()
}
