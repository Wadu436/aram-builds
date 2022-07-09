package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBConfig struct {
	ConnString string
}

var dbpool *pgxpool.Pool

func InitializeDB(config DBConfig) {
	var err error

	log.Printf("DB: Connecting to \"%v\"...\n", config.ConnString)
	dbpool, err = pgxpool.Connect(context.Background(), config.ConnString)
	if err != nil {
		log.Fatal(err)
	}
	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Fatal("DB: ERROR: Couldn't connect to database")
	}
	log.Println("DB: Connected")
}
