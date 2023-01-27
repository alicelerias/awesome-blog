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
	fmt.Println("toUint", toUint)
	fmt.Println("uid", uid)
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

// func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	user := models.User{}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	tokenID, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if tokenID != uint32(uid) {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	user.Prepare()
// 	err = user.Validate("update")
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	updatedUser, err := user.UpdateAuser(server.DB, uint32(uid))
// 	if err != nil {
// 		formattedError := formaterror.FormatError(err.Error())
// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, updatedUser)
// }

// func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)

// 	user := models.User{}

// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	tokenID, err := auth.ExtractTokenID(r)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if tokenID != 0 && tokenID != uint32(uid) {
// 		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	_, err = user.DeleteAUser(server.DB, uint32(uid))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
// 	responses.JSON(w, http.StatusNoContent, "")
// }
