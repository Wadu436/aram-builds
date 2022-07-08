package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "github.com/wadu436/gin-auth"
)

type ServerConfig struct {
	BindAddr string
	Auth     auth.Auth
}

func Server(config ServerConfig) {
	r := gin.Default()

	r.Use(config.Auth.BasicAuthMiddleware)

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Secret things :D")
		return
	})

	r.Run(config.BindAddr)
}
