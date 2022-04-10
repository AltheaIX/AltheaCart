package models

import (
	"AlthCart/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
)

var JWTKey = []byte("AlthCart")

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

func ValidateAuth(username string, password string) (User, error) {
	db := config.OpenConn()
	defer db.Close()

	user := User{}
	stmt, err := db.Preparex("SELECT id, username FROM users WHERE username=$1 and password=$2")
	if err != nil {
		return user, err
	}

	err = stmt.Get(&user, username, password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateJWT(User User) string {
	token, err := jwt.Sign(jwt.HS256, JWTKey, User, jwt.MaxAge(1*time.Hour))
	if err != nil {
		panic(err)
	}

	return string(token)
}

func VerifyJWT(token string) (User, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, JWTKey, []byte(token))
	data := User{}
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}
	err = json.Unmarshal(verifiedToken.Payload, &data)
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}
	return data, nil
}

func TestValidateAuth() {
	data, err := ValidateAuth("Malik", "Althea123")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Wrong username or password.")
			return
		}
		fmt.Println(err.Error())
	}
	fmt.Println(data)
}
