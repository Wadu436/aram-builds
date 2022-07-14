package db

import (
	"context"
	"fmt"
	"log"

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
	dbpool, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return err
	}
	err = dbpool.Ping(context.Background())
	if err != nil {
		return err
	}
	log.Println("DB: Connected")
	return nil
}
