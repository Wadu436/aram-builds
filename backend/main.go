package main

import (
	"crypto"
	"fmt"
	"log"

	"github.com/wadu436/aram-builds/backend/db"
	"github.com/wadu436/aram-builds/backend/server"
	auth "github.com/wadu436/gin-auth"
)

func main() {
	dbUser := "postgres"
	dbPass := "postgres"
	dbAddr := "localhost:5432"
	dbName := "test"

	dsn := fmt.Sprintf("postgresql://%v:%v@%v/%v", dbUser, dbPass, dbAddr, dbName)
	databaseConfig := db.DBConfig{
		ConnString: dsn,
	}

	db.InitializeDB(databaseConfig)
	err := db.Migrate()
	if err != nil {
		log.Fatal("Error during migrations. Shutting Down.")
	}

	auth := auth.Auth{
		LoadUser:   db.LoadUser,
		StoreUser:  db.StoreUser,
		Hash:       crypto.SHA256,
		SaltLength: 32,
	}

	serverConfig := server.ServerConfig{
		BindAddr: "localhost:8080",
		Auth:     auth,
	}

	auth.UpdateUser("foo", "bar")
	user, _ := auth.LoadUser("foo")
	log.Println(user)

	server.Server(serverConfig)
}
