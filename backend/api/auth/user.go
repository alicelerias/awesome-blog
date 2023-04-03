package auth

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/alicelerias/blog-golang/database"
	"github.com/alicelerias/blog-golang/models"
	"golang.org/x/crypto/bcrypt"
)

func saltPassword(salt []byte, password string) []byte {
	return append(salt, []byte(password)...)
}

func hashPassword(password string) (hash, salt []byte) {
	salt = make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	saltedPassword := saltPassword(salt, password)
	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, 10)
	if err != nil {
		panic(err)
	}
	return hashedPassword, salt
}

func CreateUser(ctx context.Context, repository database.Repository, user *models.User) error {
	return repository.CreateUser(ctx, userToModel(user))
}

func userToModel(user *models.User) *models.User {
	hash, salt := hashPassword(user.Password)
	return &models.User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: hash,
		Salt:         salt,
		Disabled:     false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
