package controllers

import (
	"net/http"

	"github.com/alicelerias/blog-golang/types"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetRecomendations(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}
	recomendations, err := server.repository.Recomendations(uid.(string))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fromModelUsers := []*types.User{}
	for _, item := range *recomendations {
		newItem := server.userFromModel(&item, uid.(string))
		fromModelUsers = append(fromModelUsers, newItem)
	}
	ctx.JSON(http.StatusOK, gin.H{"users": fromModelUsers})
}
