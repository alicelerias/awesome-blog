package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/alicelerias/blog-golang/database"
	"github.com/alicelerias/blog-golang/types"
	"golang.org/x/crypto/bcrypt"
)

type Tokens struct {
	AccessToken string `json:"access_token"`
}

func Authenticate(ctx context.Context, repository database.Repository, creds *types.Credentials) (tokens *Tokens, err error) {
	tokens = &Tokens{}
	user, err := repository.GetUser(ctx, creds.Username)
	if err != nil {
		return
	}

	if !passwordMatch(user.PasswordHash, user.Salt, creds.Password) {
		return nil, NewPasswordDoesntMatchError()
	}
	duration := time.Hour
	if creds.RememberMe {
		duration = time.Hour * 24 * 31
	}
	expiration := time.Now().Add(duration).Unix()
	tokens.AccessToken, err = GetSignedToken(strconv.Itoa(int(user.ID)), user.UserName, expiration)
	return
}

func passwordMatch(passwordHash []byte, salt []byte, password string) bool {
	saltedPassword := saltPassword(salt, password)
	err := bcrypt.CompareHashAndPassword(passwordHash, saltedPassword)
	return err == nil
}
