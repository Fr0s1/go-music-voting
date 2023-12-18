package users

import (
	"log"
	database "user-service/pkg/db/mysql"

	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func (u *User) Save() error {
	userId, _ := GetUserIdByUsername(u.Username)

	if userId != 0 {
		return errors.New("User already exists")
	} else {
		statement, err := database.Db.Prepare("INSERT INTO Users(Username, Password) VALUES (?, ?)")

		if err != nil {
			log.Fatal(err)

			return err
		}

		hashedPassword, err := hashPassword(u.Password)

		_, err = statement.Exec(u.Username, hashedPassword)

		if err != nil {
			return err
		}

		return nil
	}
}

func (u *User) Authenticate() (bool, error) {
	statement, err := database.Db.Prepare("SELECT Password from Users where Username = ?")

	row := statement.QueryRow(u.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("Wrong username or password")
		}

		return false, err
	}

	var hashedPassword string

	err = row.Scan(&hashedPassword)

	return CheckPasswordHash(u.Password, hashedPassword), nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("select ID from Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}
