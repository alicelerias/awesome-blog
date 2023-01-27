package controllers

import (
	// "encoding/json"
	// "errors"
	// "fmt"
	// "io/ioutil"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	// "strconv"

	// "github.com/alicelerias/blog-golang/api/auth"
	// "github.com/alicelerias/blog-golang/api/formaterror"

	"github.com/alicelerias/blog-golang/models"

	// "github.com/alicelerias/blog-golang/api/responses"
	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
)

func (server *Server) CreateUser(ctx *gin.Context) {
	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user.Prepare()
	if err := user.Validate(); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := server.repository.SaveUser(ctx, user)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (server *Server) GetUsers(ctx *gin.Context) {
	user := models.User{}
	users, err := server.repository.FindAllUsers(ctx, &user)
	if err != nil {
		log.Error()
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (s *Server) GetUser(ctx *gin.Context) {
	uid := ctx.Param("id")
	toUint, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("missing id"))
		return
	}
	user, err := s.repository.FindUserByID(ctx, toUint)
	fmt.Println("user", user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s *Server) DeleteUser(ctx *gin.Context) {
	uid := ctx.Param("id")
	toInt, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("missing id"))
	}

	if err := s.repository.DeleteAUser(ctx, toInt); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted user"})

}
