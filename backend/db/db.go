package db

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	auth "github.com/wadu436/gin-auth"
)

type DBConfig struct {
	ConnString string
}

var database *sqlx.DB

func InitializeDB(config DBConfig) {
	database = sqlx.MustOpen("pgx", config.ConnString)
	database.Ping()
	fmt.Println("Database Connected!")
}

func LoadUser(username string) (auth.User, bool) {
	var user auth.User
	err := database.Get(&user, "SELECT username, password, salt FROM users WHERE username = $1;", username)
	if err != nil {
		return auth.User{}, false
	}
	return user, true
}

func StoreUser(user auth.User) {
	log.Println(user)
	_, err := database.Exec(`DELETE FROM users WHERE username = $1;`, user.Username)
	if err != nil {
		log.Fatal(err)
	}
	_, err = database.Exec(`INSERT INTO users (username, password, salt) VALUES ($1, $2, $3);`, user.Username, user.Password, user.Salt)
	if err != nil {
		log.Fatal(err)
	}
}
