package albums

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	database "music-service/pkg/db/mysql"

	"music-service/pkg/users"
)

type Album struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Genre    string `json:"genre"`
	Year     int    `json:"year"`
	Uploader *users.User
}

func (album *Album) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Albums(Name, Artist, Genre, UploaderID, ReleasedYear) VALUES(?,?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(album.Name, album.Artist, album.Genre, album.Uploader.Id, album.Year)

	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

	log.Print("Album added!")

	return id
}

func GetAlbum(albumName string, artist string) (Album, error) {
	stmt, err := database.Db.Prepare("SELECT a.ID, a.Name, a.Artist, a.Genre, a.ReleasedYear, u.ID, u.Username  FROM Albums a JOIN Users u on a.UploaderID = u.ID WHERE a.Name = ? and a.Artist = ?")

	fmt.Println("Reach GetAlbum sql")
	if err != nil {
		log.Fatal(err)
	}

	rows := stmt.QueryRow(albumName, artist)

	var album Album
	var user users.User

	err = rows.Scan(&album.ID, &album.Name, &album.Artist, &album.Genre, &album.Year, &user.Id, &user.Username)

	fmt.Println("Reach GetAlbum sql 2")

	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			fmt.Errorf("No albums exist with name %s of artist %s", albumName, artist)
		}
	}

	album.Uploader = &user

	return album, nil
}

func GetAll() []Album {
	stmt, err := database.Db.Prepare("SELECT a.ID, a.Name, a.Artist, a.Genre, a.ReleasedYear, u.ID, u.Username  FROM Albums a JOIN Users u on a.UploaderID = u.ID")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query()

	defer stmt.Close()

	var albums []Album

	for rows.Next() {
		var album Album

		var user users.User

		err := rows.Scan(&album.ID, &album.Name, &album.Artist, &album.Genre, &album.Year, &user.Id, &user.Username)

		if err != nil {
			log.Fatal(err)
		}

		album.Uploader = &user

		albums = append(albums, album)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return albums
}
