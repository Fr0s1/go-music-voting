package albums

import (
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
