package controllers

import (
	"net/http"
	"strconv"

	"github.com/alicelerias/blog-golang/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateFollow(ctx *gin.Context) {
	following := &models.Following{}
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem to authenticate user"})
		return
	}
	uidToInt, _ := parseInt(uid.(string))
	following.FollowerID = uint32(uidToInt)

	followingID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	following.FollowingID = uint32(followingID)

	if err := server.repository.Follow(ctx, following); err != nil {

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
}

func (server *Server) GetFollows(ctx *gin.Context) {
	following := models.Following{}
	followings, err := server.repository.GetFollows(ctx, &following)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"followings": followings})
}

func (server *Server) Feed(ctx *gin.Context) {
	followerId, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem to authenticate user"})
		return
	}

	followerIdToString, _ := followerId.(string)

	feed, err := server.repository.Feed(ctx, followerIdToString)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"feed": &feed})
}

func (server *Server) Unfollow(ctx *gin.Context) {
	followerId, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem to authenticate user"})
		return
	}

	followerIdToString, _ := followerId.(string)

	followingId := ctx.Param("id")

	if err := server.repository.Unfollow(ctx, followerIdToString, followingId); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "user unfollowed"})
}
