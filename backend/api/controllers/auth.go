package controllers

import (
	"net/http"

	"github.com/alicelerias/blog-golang/api/auth"
	"github.com/alicelerias/blog-golang/config"
	"github.com/alicelerias/blog-golang/types"
	"github.com/gin-gonic/gin"
)

func (server *Server) Login(ctx *gin.Context) {
	configs := config.GetConfig()
	credentials := types.Credentials{}
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	token, err := auth.Authenticate(ctx, server.repository, &credentials)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	ctx.SetCookie(configs.AuthCookie, token.AccessToken, 0, "/", configs.Host, true, true)
	ctx.AbortWithStatus(http.StatusOK)
}

func (server *Server) Logout(ctx *gin.Context) {
	configs := config.GetConfig()
	ctx.SetCookie(configs.AuthCookie, "", 0, "/", configs.Host, true, true)
	ctx.AbortWithStatus(http.StatusOK)
}
