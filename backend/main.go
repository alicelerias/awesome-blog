package main

import (
	"net/http"

	"github.com/alicelerias/blog-golang/api/controllers"
	"github.com/alicelerias/blog-golang/config"
	"github.com/alicelerias/blog-golang/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	var message interface{}
	for _, err := range c.Errors {
		log.Error()
		message = err.JSON()
	}

	c.JSON(http.StatusInternalServerError, message)
}

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
	postgresRepository := database.NewPostgresDBRepository(connection)
	server := controllers.NewServer(postgresRepository)

	r := gin.Default()

	r.Use(ErrorHandler)

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	})
	r.GET("/", server.Home)

	r.POST("/user", server.CreateUser)

	r.GET("/users", server.GetUsers)

	r.GET("/users/:id", server.GetUser)

	r.Run()
}
