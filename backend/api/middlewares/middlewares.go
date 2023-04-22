package middlewares

import (
	"net/http"

	"github.com/alicelerias/blog-golang/api/auth"
	"github.com/alicelerias/blog-golang/config"
	"github.com/alicelerias/blog-golang/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CORSMiddleware() gin.HandlerFunc {
	configs := config.GetConfig()
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", configs.AllowedHosts)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	var message interface{}
	for _, err := range c.Errors {
		log.Errorf("API Error: %s \n", err.Error())
		message = err.JSON()
	}
	if message != nil {
		c.JSON(http.StatusInternalServerError, message)
	}
}

func AuthenticationMiddleware() gin.HandlerFunc {
	configs := config.GetConfig()
	return func(ctx *gin.Context) {
		token, _ := ctx.Cookie(configs.AuthCookie)
		claims, err := auth.ValidateToken(token)
		if err != nil {
			err = errors.NewForbiddenError()
			ctx.AbortWithError(http.StatusForbidden, err)

		}
		sub, ok := claims["sub"]
		if !ok {
			panic("Invalid claims")
		}

		ctx.Set("uid", sub)
		ctx.Next()

	}
}
