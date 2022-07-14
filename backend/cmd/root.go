package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wadu436/aram-builds/backend/db"
)

var dbConfig db.DBConfig = db.DBConfig{}

var rootCmd = &cobra.Command{
	Use: "aram-builds-backend",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		unset := []string{}
		if dbConfig.User == "" {
			unset = append(unset, "postgres-user")
		}
		if dbConfig.Password == "" {
			unset = append(unset, "postgres-password")
		}
		if dbConfig.Address == "" {
			unset = append(unset, "postgres-address")
		}
		if dbConfig.DatabaseName == "" {
			unset = append(unset, "postgres-database")
		}

		if len(unset) > 0 {
			return fmt.Errorf("required flag(s) \"%v\" not set", strings.Join(unset, ", "))
		}

		// Setup db
		err := db.InitializeDB(dbConfig)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dbConfig.User, "user", "U", os.Getenv("POSTGRES_USER"), "database user, can also be set by the environment variable POSTGRES_USER")
	rootCmd.PersistentFlags().StringVarP(&dbConfig.Password, "password", "P", os.Getenv("POSTGRES_PASSWORD"), "database password, can also be set by the environment variable POSTGRES_PASSWORD")
	rootCmd.PersistentFlags().StringVarP(&dbConfig.Address, "address", "a", os.Getenv("POSTGRES_ADDRESS"), "database address and port (e.g. localhost:5432), can also be set by the environment variable POSTGRES_ADDRESS")
	rootCmd.PersistentFlags().StringVarP(&dbConfig.DatabaseName, "database", "d", os.Getenv("POSTGRES_DATABASE"), "database name, can also be set by the environment variable POSTGRES_DATABASE")
}
