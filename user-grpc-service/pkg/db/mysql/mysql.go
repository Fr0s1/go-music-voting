package database

import (
	"database/sql"
	"fmt"
	"os"

	mysql "github.com/go-sql-driver/mysql"

	logging "user-grpc/pkg/logging"
)

var (
	Db     *sql.DB
	logger = logging.Log.WithFields(logging.StandardFields)
)

func InitDB() {
	dbHost := os.Getenv("DBHOST")

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:3306", dbHost),
		DBName:               "musicvoting",
		AllowNativePasswords: true,
	}

	// cfg := mysql.Config{
	// 	User:   "root",
	// 	Passwd: "hieu2203",
	// 	Net:    "tcp",
	// 	Addr:   fmt.Sprintf("%s:3306", dbHost),
	// 	DBName: "musicvoting",
	// }
	db, err := sql.Open("mysql", cfg.FormatDSN())

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
