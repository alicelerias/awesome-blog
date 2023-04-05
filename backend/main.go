package main

import (
	"net/http"

	"github.com/alicelerias/blog-golang/api/controllers"
	"github.com/alicelerias/blog-golang/api/middlewares"
	"github.com/alicelerias/blog-golang/config"
	"github.com/alicelerias/blog-golang/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var logLevelMap = map[string]log.Level{
	"ERROR":   log.ErrorLevel,
	"DEBUG":   log.DebugLevel,
	"WARNING": log.WarnLevel,
	"INFO":    log.InfoLevel,
	"FATAL":   log.FatalLevel,
}

func main() {
	configs := config.GetConfig()
	level := logLevelMap[configs.LogLevel]

	log.SetLevel(level)

	connection, err := database.GetConnection(configs)
	if err != nil {
		log.Fatal(err)
	}

	database.MigrateDB(connection)

	postgresRepository := database.NewPostgresDBRepository(connection)
	server := controllers.NewServer(postgresRepository)

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	r.Use(middlewares.ErrorHandler)

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	})
	r.GET("/", server.Home)

	r.POST("/users", server.CreateUser)

	r.POST("/login", server.Login)

	r.POST("/logout", server.Logout)

	r.Use(middlewares.AuthenticationMiddleware())

	r.GET("/users", server.GetUsers)

	r.GET("/users/:id", server.GetUser)

	r.PUT("/users/:id", server.UpdateUser)

	r.DELETE("/users/:id", server.DeleteUser)

	r.POST("/posts", server.CreatePost)

	r.GET("/posts", server.GetPosts)

	r.GET("/posts/:id", server.GetPost)

	r.PUT("/posts/:id", server.UpdatePost)

	r.DELETE("/posts/:id", server.DeletePost)

	r.POST("/follow/:id", server.CreateFollow)

	r.GET("/follows", server.GetFollows)

	r.DELETE("/follow/:id", server.Unfollow)

	r.GET("/users/followings", server.Feed)

	r.Run()
}
