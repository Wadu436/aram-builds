package db

import (
	"context"
	"embed"
	"log"
)

var MIGRATIONS []string = []string{"000_create_users.sql", "001_create_builds.sql", "002_add_comments_field.sql", "003_add_comments_for_items.sql", "004_allow_null_arrays.sql"}

//go:embed migrations/*
var f embed.FS

var SQL_CREATE_MIGRATION = `CREATE TABLE IF NOT EXISTS migrations (migration VARCHAR(64) PRIMARY KEY);`

var SQL_ADD_MIGRATION = `INSERT INTO migrations VALUES ($1);`
var SQL_CHECK_MIGRATION = `SELECT COUNT(1) FROM migrations WHERE migration = $1;`

func Migrate() error {
	log.Printf("DB: Starting migrations\n")
	_, _ = dbpool.Exec(context.Background(), SQL_CREATE_MIGRATION)

	for _, migration := range MIGRATIONS {
		// Check if migration already applied
		var count int
		if err := dbpool.QueryRow(context.Background(), SQL_CHECK_MIGRATION, migration).Scan(&count); err != nil {
			return err
		} else if count != 0 {
			continue
		}

		// Apply migration
		m, err := f.ReadFile("migrations/" + migration)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("DB: Applying migration: %v\n", migration)

		tx, err := dbpool.Begin(context.Background())
		if err != nil {
			return err
		}
		defer tx.Rollback(context.Background())
		if _, err = tx.Exec(context.Background(), string(m)); err != nil {
			return err
		}
		if _, err = tx.Exec(context.Background(), SQL_ADD_MIGRATION, migration); err != nil {
			return err
		}
		tx.Commit(context.Background())
	}
	log.Printf("DB: Finished migrations\n")
	return nil
}
