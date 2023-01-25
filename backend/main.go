package main

import (
	"log"
	"net/http"

	"github.com/alicelerias/blog-golang/api/controllers"
	"github.com/alicelerias/blog-golang/config"
	"github.com/alicelerias/blog-golang/database"
	"github.com/gin-gonic/gin"
)

func main() {
	configs := config.GetConfig()
	connection, err := database.GetConnection(configs)
	if err != nil {
		log.Fatal(err)
	}
	postgresRepository := database.NewPostgresDBRepository(connection)
	server := controllers.NewServer(postgresRepository)

	r := gin.Default()
	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	})
	r.GET("/", server.Home)

	r.POST("/create", server.CreateUser)

	r.Run()
}
