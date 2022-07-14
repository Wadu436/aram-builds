package cmd

import (
	"crypto"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wadu436/aram-builds/backend/db"
	"github.com/wadu436/aram-builds/backend/server"
	auth "github.com/wadu436/gin-auth"
)

var serverConfig server.ServerConfig = server.ServerConfig{
	Auth: auth.Auth{
		LoadUser:   db.LoadUser,
		StoreUser:  db.StoreUser,
		Hash:       crypto.SHA256,
		SaltLength: 32,
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&serverConfig.BindAddr, "bind-address", "b", os.Getenv("BIND_ADDRESS"), "IP Address the server binds to. Can also be set by the environment variable BIND_ADDRESS.")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the REST server",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		unset := []string{}
		if serverConfig.BindAddr == "" {
			unset = append(unset, "bind-address")
		}
		if len(unset) > 0 {
			return fmt.Errorf("required flag(s) \"%v\" not set", strings.Join(unset, ", "))
		} else {
			return nil
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		server.Server(serverConfig)
	},
}
