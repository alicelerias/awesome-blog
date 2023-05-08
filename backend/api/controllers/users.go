package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"net/http"

	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"

	"github.com/alicelerias/blog-golang/api/auth"
	"github.com/alicelerias/blog-golang/models"
	"github.com/alicelerias/blog-golang/types"
	"github.com/gin-gonic/gin"

	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

func (server *Server) userFromModel(model *models.User, followerId string) *types.User {
	followingId := strconv.Itoa(int(model.ID))
	return &types.User{
		ID:          model.ID,
		UserName:    model.UserName,
		Email:       model.Email,
		Bio:         model.Bio,
		Avatar:      model.Avatar,
		IsFollowing: server.repository.IsFollowing(followerId, followingId),
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func (server *Server) CreateUser(ctx *gin.Context) {
	user := &models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "translator not found"})
		return
	}
	v := validator.New()

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error on register default translations"})
		return
	}
	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("passwd", trans, func(ut ut.Translator) error {
		return ut.Add("passwd", "{0} is not strong enough", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("passwd", fe.Field())
		return t
	})

	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})

	if err := v.Struct(user); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, e.Translate(trans))
		}
		return

	}

	if err := auth.CreateUser(server.repository, user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.AbortWithStatus(http.StatusCreated)
	if err := server.cache.DeleteKey("users", "all"); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
}

func (server *Server) GetUsers(ctx *gin.Context) {
	user := models.User{}
	cache := []*types.User{}
	name := "users"
	nameSpace := "all"
	if err := server.cache.GetKey(name, nameSpace, &cache); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"users": cache})
	} else {
		users, err := server.repository.FindAllUsers(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		id, exists := ctx.Get("uid")
		if !exists {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		idToString := id.(string)

		fromModelUsers := []*types.User{}
		for _, item := range *users {
			newItem := server.userFromModel(&item, idToString)
			fromModelUsers = append(fromModelUsers, newItem)
		}
		ctx.JSON(http.StatusOK, gin.H{"users": fromModelUsers})
		server.cache.SetKey(name, nameSpace, fromModelUsers, time.Hour)
	}

}

func (s *Server) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := s.repository.FindUserByID(id)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	followerId, exists := ctx.Get("uid")
	if !exists {
		ctx.AbortWithError(http.StatusForbidden, err)
		return
	}

	followerIdToString := followerId.(string)
	userFromModel := s.userFromModel(user, followerIdToString)
	ctx.JSON(http.StatusOK, userFromModel)
}

func (s *Server) GetCurrentUser(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}
	key := "user_profile"
	cache := types.User{}

	err := s.cache.GetKey(key, uid.(string), &cache)

	if err == nil {
		ctx.JSON(http.StatusOK, cache)
	} else {
		user, err := s.repository.FindUserByID(uid.(string))
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		userFromModel := s.userFromModel(user, uid.(string))
		err = s.cache.SetKey(key, uid.(string), userFromModel, time.Hour)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, userFromModel)
	}

}

func (s *Server) UpdateUser(ctx *gin.Context) {
	userId, exists := ctx.Get("uid")
	if !exists {
		ctx.AbortWithError(http.StatusForbidden, errors.New("Error on authenticate user!"))
		return
	}

	userIdToString := userId.(string)

	whiteList := []string{"bio", "avatar"}
	input, err := getValidJson(ctx.Request.Body, whiteList)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid data!"))
		return
	}
	input["updated_at"] = time.Now()
	user, err := s.repository.UpdateUser(input, userIdToString)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	userFromModel := s.userFromModel(user, userIdToString)

	ctx.JSON(http.StatusOK, userFromModel)

}

func (s *Server) UpdateCurrentUser(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "problem to authenticate user"})
		return
	}
	whiteList := []string{"bio", "avatar"}
	input, err := getValidJson(ctx.Request.Body, whiteList)
	if err != nil {
		fmt.Println(input)
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid data!"))
		return
	}
	input["updated_at"] = time.Now()
	user, err := s.repository.UpdateUser(input, uid.(string))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	userFromModel := s.userFromModel(user, uid.(string))
	ctx.JSON(http.StatusOK, userFromModel)
	err = s.cache.DeleteKey("user_profile", uid.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
}

func (s *Server) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := s.repository.DeleteUser(id); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "successfully deleted user"})

}
