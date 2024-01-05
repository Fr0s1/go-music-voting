package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"

	logging "user-service/pkg/logging"
)

var Db *sql.DB

func InitDB() {

	logger := logging.Log.WithFields(logging.StandardFields)

	dbHost := os.Getenv("DBHOST")

	// cfg := mysql.Config{
	// 	User:   os.Getenv("DBUSER"),
	// 	Passwd: os.Getenv("DBPASS"),
	// 	Net:    "tcp",
	// 	Addr:   fmt.Sprintf("%s:3306", dbHost),
	// 	DBName: "musicvoting",
	// }

	cfg := mysql.Config{
		User:   "root",
		Passwd: "hieu2203",
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:3306", dbHost),
		DBName: "musicvoting",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		logger.Error(err)
	}

	if err = db.Ping(); err != nil {
		logger.Error(err)
	}

	Db = db

	logger.Info("Connected")
}

func CloseDB() error {
	return Db.Close()
}
