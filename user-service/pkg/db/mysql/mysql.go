package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:hieu2203@tcp(localhost)/musicvoting")

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
