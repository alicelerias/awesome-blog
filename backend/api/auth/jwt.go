package auth

import (
	"fmt"
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

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("error 1")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		fmt.Println("error 2")
		return err
	}

	if !token.Valid {
		fmt.Println("error 3")
		return NewInvalidToken("Invalid token!")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			fmt.Println("error 4")
			return NewInvalidToken("Invalid exp value")
		}
		if time.Now().Unix() > int64(exp) {
			fmt.Println("error 5")
			return NewInvalidToken("Token expired")
		}
	}
	return nil
}
