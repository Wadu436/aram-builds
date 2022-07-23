package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wadu436/aram-builds/backend/db"
	auth "github.com/wadu436/gin-auth"
)

type ServerConfig struct {
	BindAddr       string
	TrustedProxies []string
	Auth           auth.Auth
}

func Server(config ServerConfig) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.SetTrustedProxies(config.TrustedProxies)

	// API group
	api := r.Group("/api")

	// Auth routes
	auth := api.Group("/", config.Auth.BasicAuthMiddleware)
	auth.POST("auth/check/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Build routes
	build := api.Group("/build")
	buildAuth := build.Group("/", config.Auth.BasicAuthMiddleware)

	build.GET("/", func(c *gin.Context) {
		builds, err := db.AllBuilds()
		if err == pgx.ErrNoRows {
			ErrorNotFound(c)
			return
		}
		if err != nil {
			log.Println(err)
			ErrorDatabase(c)
			return
		}
		c.JSON(http.StatusOK, builds)
	})

	buildAuth.POST("/", func(c *gin.Context) {
		var build db.Build
		err := c.ShouldBindJSON(&build)
		if err != nil {
			ErrorBadRequest(c)
			log.Println(err)
			return
		}
		build.Mtime = db.JSONTime(time.Now().UTC())

		db.StoreBuild(build)
		c.JSON(http.StatusOK, gin.H{"mtime": build.Mtime})
	})

	build.GET("/:champion/", func(c *gin.Context) {
		builds, err := db.AllBuildsChampion(c.Param("champion"))
		if err == pgx.ErrNoRows {
			ErrorNotFound(c)
			return
		}
		if err != nil {
			log.Println(err)
			ErrorDatabase(c)
			return
		}
		c.JSON(http.StatusOK, builds)
	})

	build.GET("/:champion/:version", func(c *gin.Context) {
		champion := c.Param("champion")
		version := c.Param("version")

		build, err := db.LoadBuild(champion, version)
		if err == pgx.ErrNoRows {
			ErrorNotFound(c)
			return
		}
		if err != nil {
			ErrorDatabase(c)
			return
		}

		c.JSON(http.StatusOK, build)
	})

	buildAuth.DELETE("/:champion/:version", func(c *gin.Context) {
		champion := c.Param("champion")
		version := c.Param("version")
		err := db.DeleteBuild(champion, version)
		if err != nil {
			ErrorDatabase(c)
		}
	})

	r.Run(config.BindAddr)
}

// Errors
func ErrorDatabase(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Database error"})
}

func ErrorBadRequest(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": 2, "message": "Bad Request"})

}

func ErrorNotFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": 3, "message": "Not Found"})
}
