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

	auth := r.Group("/", config.Auth.BasicAuthMiddleware)

	auth.POST("auth/check/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.Run(config.BindAddr)
}
