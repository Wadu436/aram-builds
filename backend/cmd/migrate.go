package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/wadu436/aram-builds/backend/db"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations on the database",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Migrate()
		if err != nil {
			log.Fatalf("Error during migrations: %v\n", err)
		}
	},
}
