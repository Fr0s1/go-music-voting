package models

import (
	"database/sql"
	"errors"
	"fmt"
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

func AddAlbumPoll(albumId int64, pollId int64) int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Poll_Album(AlbumID, PollID) VALUES (?,?)")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	res, err := stmt.Exec(albumId, pollId)

	if err != nil {
		fmt.Println(err.Error())
	}

	id, _ := res.LastInsertId()

	return id
}

func GetAlbumDetails(albumId int64) Album {
	stmt, err := database.Db.Prepare("SELECT Id, Name, Artist FROM Albums WHERE ID = (?)")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	rows := stmt.QueryRow(albumId)

	var album Album

	err = rows.Scan(&album.Id, &album.Name, &album.Artist)

	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			fmt.Errorf("No albums exist with ID %d", albumId)
		}
	}

	return album
}

func VotePollAlbum(pollId int64, albumId int64, voterId int32) (int64, error) {
	stmt, err := database.Db.Prepare("INSERT INTO Votes(AlbumID, PollID, VoterID) VALUES(?,?,?)")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer stmt.Close()

	res, err := stmt.Exec(albumId, pollId, voterId)

	if err != nil {
		fmt.Println(err.Error())
	}

	id, err := res.LastInsertId()

	return id, nil
}
