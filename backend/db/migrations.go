package db

import (
	"embed"
	"log"
)

var MIGRATIONS []string = []string{"000_create_users.sql"}

//go:embed migrations/*
var f embed.FS

var SQL_CREATE_MIGRATION = `CREATE TABLE IF NOT EXISTS migrations (migration VARCHAR(64) PRIMARY KEY);`

var SQL_ADD_MIGRATION = `INSERT INTO migrations VALUES ($1);`
var SQL_CHECK_MIGRATION = `SELECT COUNT(1) FROM migrations WHERE migration = $1;`

func Migrate() {
	log.Printf("DB: Starting migrations\n")
	database.MustExec(SQL_CREATE_MIGRATION)

	for _, migration := range MIGRATIONS {
		// Check if migration already applied
		var count int
		if database.QueryRow(SQL_CHECK_MIGRATION, migration).Scan(&count); count != 0 {
			continue
		}

		// Apply migration
		m, err := f.ReadFile("migrations/" + migration)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("DB: Applying migration: %v\n", migration)

		database.MustExec(string(m))
		database.MustExec(SQL_ADD_MIGRATION, migration)
	}
	log.Printf("DB: Finished migrations\n")
}
