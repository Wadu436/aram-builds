package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBConfig struct {
	User         string
	Password     string
	Address      string
	DatabaseName string
}

var dbpool *pgxpool.Pool

func InitializeDB(config DBConfig) error {
	var err error
	connString := fmt.Sprintf("postgresql://%v:%v@%v/%v", config.User, config.Password, config.Address, config.DatabaseName)

	log.Printf("DB: Connecting to \"%v\"...\n", connString)
	for retries := 5; retries > 0; retries-- {
		dbpool, err = pgxpool.Connect(context.Background(), connString)
		if err == nil {
			break
		} else {
			log.Printf("DB: Error connecting to database: \"%v\". Retrying...\n", err)
			time.Sleep(time.Second)
		}
	}
	log.Println("DB: Connected")
	return nil
}
