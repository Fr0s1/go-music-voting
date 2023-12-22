package models

import (
	"log"
	database "voting-grpc/pkg/db/mysql"
)

type User struct {
	Id       int32
	Username string
}

type Poll struct {
	Name    string
	Creator *User
	Albums  []*Album
}

type Album struct {
	Name   string
	Artist string
	Id     int64
}

type Vote struct {
	Voter *User
	Poll  *Poll
	Album *Album
}

func (poll *Poll) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Polls(Name, CreatorID) VALUES (?,?)")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(poll.Name, poll.Creator.Id)

	if err != nil {
		log.Fatal(err)
	}

	id, _ := res.LastInsertId()

	return id
}
