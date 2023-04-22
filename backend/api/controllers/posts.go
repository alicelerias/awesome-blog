package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/alicelerias/blog-golang/models"
	"github.com/alicelerias/blog-golang/types"
	"github.com/gin-gonic/gin"
)

func NewUser(user *models.User) *types.User {
	return &types.User{
		ID:       user.ID,
		UserName: user.UserName,
	}
}

func (server *Server) postFromModel(ctx *gin.Context, post *models.Post, user *models.User, userId string) *types.Post {
	postId := strconv.Itoa(int(post.ID))
	return &types.Post{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		Img:        post.Img,
		Author:     *NewUser(user),
		AuthorID:   post.AuthorID,
		IsFavorite: server.repository.GetFavorite(ctx, postId, userId),
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}
}

func (server *Server) CreatePost(ctx *gin.Context) {
	post := &models.Post{}
	if err := ctx.ShouldBind(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem to authenticate user"})
		return
	}
	postUid, _ := strconv.ParseUint(uid.(string), 10, 64)
	post.AuthorID = uint32(postUid)

	if err := server.repository.CreatePost(ctx, post); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
}

func (server *Server) GetPosts(ctx *gin.Context) {
	post := models.Post{}
	posts, err := server.repository.GetPosts(ctx, &post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fromModelPosts := []*types.Post{}

	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem to authenticate user"})
		return
	}

	for _, item := range *posts {
		newPost := server.postFromModel(ctx, &item, &item.Author, uid.(string))
		fromModelPosts = append(fromModelPosts, newPost)
	}

	ctx.JSON(http.StatusOK, gin.H{"feed": fromModelPosts})
}
func (server *Server) GetBlogPosts(ctx *gin.Context) {
	post := *&models.Post{}
	cursor := ctx.Query("cursor")
	limit := 10

	id := ctx.Param("id")

	feed, err := server.repository.GetPostsByUser(ctx, &post, cursor, id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fromModelPosts := []*types.Post{}

	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "problem to authenticate user"})
		return
	}

	for _, item := range feed {
		newPost := server.postFromModel(ctx, &item, &item.Author, uid.(string))
		fromModelPosts = append(fromModelPosts, newPost)
	}

	if len(feed) == limit {
		nextCursor := feed[len(feed)-1].CreatedAt

		nextLink := fmt.Sprintf("/feed?cursor=%s", url.QueryEscape(nextCursor.Format(time.RFC3339Nano)))

		ctx.JSON(http.StatusOK, gin.H{
			"feed":        feed,
			"next_cursor": nextCursor.Format(time.RFC3339),
			"next_link":   nextLink,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"feed": fromModelPosts})
	}
}

func (server *Server) GetPostsByUser(ctx *gin.Context) {
	post := *&models.Post{}
	cursor := ctx.Query("cursor")
	limit := 10

	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}

	uidToString, _ := uid.(string)

	feed, err := server.repository.GetPostsByUser(ctx, &post, cursor, uidToString)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fromModelPosts := []*types.Post{}

	for _, item := range feed {
		newPost := server.postFromModel(ctx, &item, &item.Author, uid.(string))
		fromModelPosts = append(fromModelPosts, newPost)
	}

	if len(feed) == limit {
		nextCursor := feed[len(feed)-1].CreatedAt

		nextLink := fmt.Sprintf("/feed?cursor=%s", url.QueryEscape(nextCursor.Format(time.RFC3339Nano)))

		ctx.JSON(http.StatusOK, gin.H{
			"feed":        feed,
			"next_cursor": nextCursor.Format(time.RFC3339),
			"next_link":   nextLink,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"feed": fromModelPosts})
	}
}

func (server *Server) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := server.repository.GetPost(ctx, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}

	uidToString, _ := uid.(string)

	ctx.JSON(http.StatusOK, server.postFromModel(ctx, post, &post.Author, uidToString))
}

func (server *Server) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	whiteList := []string{"title", "content", "img"}
	input, err := getValidJson(ctx.Request.Body, whiteList)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid data!"))
		return
	}
	input["updated_at"] = time.Now()

	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}

	uidToString, _ := uid.(string)

	post, err := server.repository.UpdatePost(ctx, input, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, server.postFromModel(ctx, post, &post.Author, uidToString))
}

func (server *Server) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := server.repository.DeletePost(ctx, id); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "successfully deleted post"})
}
