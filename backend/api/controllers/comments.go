package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/alicelerias/blog-golang/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateComment(ctx *gin.Context) {
	comment := &models.Comment{}
	if err := ctx.ShouldBind(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Problem to authenticate user"})
		return
	}
	commentAuthorId, _ := strconv.ParseUint(uid.(string), 10, 64)
	comment.AuthorId = uint32(commentAuthorId)

	postId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	comment.PostId = uint32(postId)

	if err := server.repository.CreateComment(ctx, comment); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
}

func (server *Server) DeleteComment(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	userId, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}

	userIdToUint, _ := parseInt(userId.(string))

	if err := server.repository.DeleteComment(ctx, uint32(id), uint32(userIdToUint)); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}

func (server *Server) GetPostComments(ctx *gin.Context) {
	postId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	cursor := ctx.Query("cursor")
	limit := 10
	comments, err := server.repository.GetPostComments(ctx, cursor, uint32(postId))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(comments) == limit {
		nextCursor := comments[len(comments)-1].CreatedAt

		nextLink := fmt.Sprintf("/feed?cursor=%s", url.QueryEscape(nextCursor.Format(time.RFC3339Nano)))

		ctx.JSON(http.StatusOK, gin.H{
			"comments":    comments,
			"next_cursor": nextCursor.Format(time.RFC3339),
			"next_link":   nextLink,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"comments": &comments})
}
