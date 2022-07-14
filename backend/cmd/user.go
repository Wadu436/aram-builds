package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(userCmd)

	userCmd.Flags().StringVarP(&user, "new-username", "u", "", "Username for the new user")
	userCmd.Flags().StringVarP(&user, "new-password", "p", "", "Password for the new user")

	userCmd.MarkFlagRequired("new-username")
	userCmd.MarkFlagRequired("new-password")
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User management",
	Run: func(cmd *cobra.Command, args []string) {
		serverConfig.Auth.UpdateUser(user, password)
		log.Printf("Updated user \"%v\"", user)
	},
}

var user, password string
