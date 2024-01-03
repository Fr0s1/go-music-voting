// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Album struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Artist   string `json:"artist"`
	Genre    string `json:"genre"`
	Uploader *User  `json:"uploader"`
}

type AlbumSearch struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

type NewAlbum struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}

type NewPoll struct {
	Name   string      `json:"name"`
	Albums []*NewAlbum `json:"albums"`
}

type NewVote struct {
	PollID  string `json:"pollId"`
	AlbumID string `json:"albumId"`
}

type Poll struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	AlbumVotes []*PollAlbum `json:"albumVotes"`
}

type PollAlbum struct {
	Poll  *Poll   `json:"poll"`
	Votes []*Vote `json:"votes"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Vote struct {
	PollID  string `json:"pollId"`
	AlbumID string `json:"albumId"`
}
