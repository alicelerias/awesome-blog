package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/alicelerias/blog-golang/models"
	"github.com/alicelerias/blog-golang/types"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateFollow(ctx *gin.Context) {
	following := &models.Following{}
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}
	uidToInt, _ := parseInt(uid.(string))
	following.FollowerID = uint32(uidToInt)

	followingID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	following.FollowingID = uint32(followingID)

	if err := server.repository.Follow(following); err != nil {

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
}

func (server *Server) GetFollows(ctx *gin.Context) {
	following := models.Following{}
	followings, err := server.repository.GetFollows(&following)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"followings": followings})
}

func (server *Server) Feed(ctx *gin.Context) {
	cursor := ctx.Query("cursor")
	limit := 10

	followerId, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}

	followerIdString, _ := followerId.(string)

	feed, err := server.repository.Feed(cursor, followerIdString)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fromModelFeed := []*types.Post{}

	for _, item := range feed {
		newPost := server.postFromModel(ctx, &item, &item.Author, followerIdString)
		fromModelFeed = append(fromModelFeed, newPost)
	}
	if len(feed) == limit {
		nextCursor := feed[len(feed)-1].CreatedAt

		nextLink := fmt.Sprintf("/feed?cursor=%s", url.QueryEscape(nextCursor.Format(time.RFC3339Nano)))

		ctx.JSON(http.StatusOK, gin.H{
			"content":     fromModelFeed,
			"next_cursor": nextCursor.Format(time.RFC3339),
			"next_link":   nextLink,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"content": fromModelFeed})
	}
}

func (server *Server) Unfollow(ctx *gin.Context) {
	followerId, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}

	followerIdToString, _ := followerId.(string)

	followingId := ctx.Param("id")

	if err := server.repository.Unfollow(followerIdToString, followingId); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "user unfollowed"})
}
