package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	dbHost := os.Getenv("DBHOST")

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:3306", dbHost),
		DBName: "musicvoting",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	Db = db

	fmt.Println("Connected")
}

func CloseDB() error {
	return Db.Close()
}
