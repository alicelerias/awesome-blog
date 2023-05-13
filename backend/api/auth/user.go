package auth

import (
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

func CreateUser(repository database.Repository, user *models.User) error {
	userToModel(user)
	return repository.CreateUser(user)
}

func userToModel(user *models.User) {
	hash, salt := hashPassword(user.Password)
	user.Salt = salt
	user.PasswordHash = hash
	user.Disabled = false
	user.Password = ""
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

}
