package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wadu436/aram-builds/backend/db"
	auth "github.com/wadu436/gin-auth"
)

type ServerConfig struct {
	BindAddr string
	Auth     auth.Auth
}

func Server(config ServerConfig) {
	r := gin.Default()

	// Auth routes
	auth := r.Group("/", config.Auth.BasicAuthMiddleware)
	auth.POST("auth/check/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Build routes
	build := r.Group("/build")
	buildAuth := build.Group("/", config.Auth.BasicAuthMiddleware)

	build.GET("/", func(c *gin.Context) {
		builds, err := db.AllBuilds()
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
		if err != nil {
			log.Println(err)
			ErrorDatabase(c)
			return
		}
		c.JSON(http.StatusOK, builds)
	})

	build.GET("/:champion/latest", func(c *gin.Context) {
		build, exists := db.LatestBuild(c.Param("champion"))
		if !exists {
			ErrorNotFound(c)
			return
		}
		c.JSON(http.StatusOK, build)
	})

	build.GET("/:champion/:major/:minor", func(c *gin.Context) {
		champion := c.Param("champion")
		major, err1 := strconv.Atoi(c.Param("major"))
		minor, err2 := strconv.Atoi(c.Param("minor"))
		if err1 != nil || err2 != nil {
			ErrorBadRequest(c)
		}

		build, exists := db.LoadBuild(champion, major, minor)
		if !exists {
			ErrorNotFound(c)
			return
		}
		c.JSON(http.StatusOK, build)
	})

	buildAuth.DELETE("/:champion/:major/:minor", func(c *gin.Context) {
		champion := c.Param("champion")
		major, err1 := strconv.Atoi(c.Param("major"))
		minor, err2 := strconv.Atoi(c.Param("minor"))
		if err1 != nil || err2 != nil {
			ErrorBadRequest(c)
		}
		db.DeleteBuild(champion, major, minor)
	})

	r.Run(config.BindAddr)
}

// Errors
func ErrorDatabase(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": 1, "message": "Database error"})
}

func ErrorBadRequest(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": 2, "message": "Bad Request"})

}

func ErrorNotFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": 3, "message": "Not Found"})

}
