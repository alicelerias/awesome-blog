package controllers

import (
	"errors"
	"fmt"

	"net/http"

	"time"

	log "github.com/sirupsen/logrus"

	"github.com/alicelerias/blog-golang/api/auth"
	"github.com/alicelerias/blog-golang/models"
	"github.com/alicelerias/blog-golang/types"
	"github.com/gin-gonic/gin"
)

func userFromModel(model *models.User) *types.User {
	return &types.User{
		ID:        model.ID,
		UserName:  model.UserName,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (server *Server) CreateUser(ctx *gin.Context) {
	user := &models.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := auth.CreateUser(ctx, server.repository, user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
}

func (server *Server) GetUsers(ctx *gin.Context) {
	user := models.User{}
	users, err := server.repository.FindAllUsers(ctx, &user)
	if err != nil {
		log.Error()
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fromModelUsers := []*types.User{}
	for _, item := range *users {
		newItem := userFromModel(&item)
		fromModelUsers = append(fromModelUsers, newItem)
	}
	ctx.JSON(http.StatusOK, gin.H{"users": fromModelUsers})
}

func (s *Server) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := s.repository.FindUserByID(ctx, id)
	fmt.Println("user", user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, userFromModel(user))
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	whiteList := []string{"user_name"}
	input, err := getValidJson(ctx.Request.Body, whiteList)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid data!"))
		return
	}
	input["updated_at"] = time.Now()
	user, err := s.repository.UpdateUser(ctx, input, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, userFromModel(user))
}

func (s *Server) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := s.repository.DeleteUser(ctx, id); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted user"})

}
