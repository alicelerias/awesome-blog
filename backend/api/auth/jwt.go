package auth

import (
	"time"

	"github.com/alicelerias/blog-golang/config"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetSignedToken(sub string, username string, exp int64) (token string, err error) {
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      sub,
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      exp,
	})

	token, err = unsignedToken.SignedString(config.GetConfig().JWTSecret)
	return
}
