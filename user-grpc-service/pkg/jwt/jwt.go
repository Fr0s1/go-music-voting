package jwt

import (
	"database/sql"
	"log"

	"github.com/dgrijalva/jwt-go"

	database "user-grpc/pkg/db/mysql"

	user "user-grpc/pkg/grpc"
)

var (
	SecretKey = []byte("secret")
)

func ParseToken(tokenStr string) (user.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)

		statement, err := database.Db.Prepare("select ID, Username from Users WHERE Username = ?")
		if err != nil {
			log.Fatal(err)
		}
		row := statement.QueryRow(username)

		var Id int32
		var Username string
		err = row.Scan(&Id, &Username)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Print(err)
			}
			return user.User{}, err
		}

		return user.User{Id: Id, Username: Username}, nil
	} else {
		return user.User{}, err
	}
}
