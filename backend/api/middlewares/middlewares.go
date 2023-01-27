package middlewares

import (
	"net/http"

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
