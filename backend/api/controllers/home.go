package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Home(ctx *gin.Context) {
	s.repository.GetHome()
	ctx.JSON(http.StatusOK, "Hello, World!")
}
