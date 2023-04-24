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

func (server *Server) Favorite(ctx *gin.Context) {
	favorite := &models.Favorite{}
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem to authenticate user"})
		return
	}
	uidToInt, _ := parseInt(uid.(string))
	favorite.UserId = uint32(uidToInt)

	postId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	favorite.PostId = uint32(postId)

	if err := server.repository.Favorite(ctx, favorite); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
}

func (server *Server) Unfavorite(ctx *gin.Context) {
	postId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	userId, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem to authenticate user"})
		return
	}

	userIdToUint, _ := parseInt(userId.(string))

	if err := server.repository.Unfavorite(ctx, uint32(postId), uint32(userIdToUint)); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}

func (server *Server) GetFavoritesByPost(ctx *gin.Context) {
	postId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	favorites, err := server.repository.GetFavoritesByPost(ctx, uint32(postId))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"feed": &favorites})
}

func (server *Server) GetFavoritesPosts(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	limit := 10
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem to authenticate user"})
		return
	}

	uidToInt, _ := parseInt(uid.(string))

	posts, err := server.repository.GetFavoritesPostsByUser(ctx, uint32(uidToInt))

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fromModelPosts := []*types.Post{}

	for _, item := range *posts {
		newPost := server.postFromModel(ctx, &item, &item.Author, uid.(string))
		fromModelPosts = append(fromModelPosts, newPost)
	}
	if len(*posts) == limit {
		nextCursor := fromModelPosts[len(fromModelPosts)-1].CreatedAt

		nextLink := fmt.Sprintf("/feed?cursor=%s", url.QueryEscape(nextCursor.Format(time.RFC3339Nano)))

		ctx.JSON(http.StatusOK, gin.H{
			"feed":        fromModelPosts,
			"next_cursor": nextCursor.Format(time.RFC3339),
			"next_link":   nextLink,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"feed": fromModelPosts})
	}
}
