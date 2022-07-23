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
		LoadUser: func(username string) (user auth.User, exists bool) {
			exists = true
			user, err := db.LoadUser(username)
			if err != nil {
				exists = false
			}
			return
		},
		StoreUser:  db.StoreUser,
		Hash:       crypto.SHA256,
		SaltLength: 32,
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&serverConfig.BindAddr, "bind-address", "b", os.Getenv("BIND_ADDRESS"), "ip address the server binds to, can also be set by the environment variable BIND_ADDRESS")
	serverCmd.Flags().StringSliceVarP(&serverConfig.TrustedProxies, "trusted-proxies", "t", strings.Split(os.Getenv("TRUSTED_PROXIES"), ","), "which proxies the server should trust, splits on commas, can also be set by the environment variable BIND_ADDRESS, which also splits on commas (e.g. \"127.0.0.1:80,192.168.2.1\")")
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
