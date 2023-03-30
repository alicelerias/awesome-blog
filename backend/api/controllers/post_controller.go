package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/alicelerias/blog-golang/models"
	"github.com/gin-gonic/gin"
)

func NewUser(user *models.User) *models.User {
	return &models.User{
		ID:       user.ID,
		UserName: user.UserName,
	}
}

func postFromModel(post *models.Post, user *models.User) *models.Post {
	return &models.Post{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Img:       post.Img,
		Author:    *NewUser(user),
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func (server *Server) CreatePost(ctx *gin.Context) {
	post := &models.Post{}
	if err := ctx.ShouldBind(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	uid, _ := ctx.Get("uid")
	fmt.Println("UUUUUID", uid)
	postId, _ := strconv.ParseUint(uid.(string), 10, 64)
	post.AuthorID = uint32(postId)

	if err := server.repository.CreatePost(ctx, post); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
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
	fromModelPosts := []*models.Post{}

	for _, item := range *posts {
		newPost := postFromModel(&item, &item.Author)
		fromModelPosts = append(fromModelPosts, newPost)
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": fromModelPosts})
}

func (server *Server) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := server.repository.GetPost(ctx, id)
	fmt.Println("POST", post)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, postFromModel(post, &post.Author))
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
	post, err := server.repository.UpdatePost(ctx, input, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, postFromModel(post, &post.Author))
}

func (server *Server) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := server.repository.DeletePost(ctx, id); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted post"})
}
