package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alicelerias/blog-golang/api/auth"
	"github.com/alicelerias/blog-golang/config"
	"github.com/alicelerias/blog-golang/errors"
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
	if message != nil {
		c.JSON(http.StatusInternalServerError, message)
	}
}

func AuthenticationMiddleware() gin.HandlerFunc {
	configs := config.GetConfig()
	return func(ctx *gin.Context) {
		token, _ := ctx.Cookie(configs.AuthCookie)
		fmt.Println("TOKEN:" + token)
		if err := auth.ValidateToken(token); err != nil {
			err = errors.NewForbiddenError()
			ctx.AbortWithError(http.StatusForbidden, err)
		}
		ctx.Next()
	}
}
