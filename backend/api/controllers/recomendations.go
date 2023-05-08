package controllers

import (
	"net/http"

	"github.com/alicelerias/blog-golang/types"
	"github.com/gin-gonic/gin"
)

func (server *Server) GetRecomendations(ctx *gin.Context) {
	// cache := []*types.User{}
	// name := "users"
	// nameSpace := "all"
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}
	// if err := server.cache.GetKey(name, nameSpace, &cache); err == nil {
	// 	ctx.JSON(http.StatusOK, gin.H{"users": cache})
	// } else {
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
	// server.cache.SetKey(name, nameSpace, fromModelUsers, time.Hour)
	// }
}
