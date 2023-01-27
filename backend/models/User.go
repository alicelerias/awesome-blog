package models

import (
	"crypto/rand"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName       string    `gorm:"size:255;not null;unique" json:"username" validate:"required,min=4,max=15"`
	Email          string    `gorm:"size:100;not null;unique" json:"email" validate:"required,email"`
	Password       string    `gorm:"size:100;not null;" json:"password" validate:"required,min=6,max=15"`
	PasswordHash   []byte    `gorm:"not null" json:"pwd_hash"`
	Salt           []byte    `gorm:"not null" json:"salt"`
	Disabled       bool      `gorm:"not null" json:"disabled"`
	DisabledReason string    `json:"disabled_reason"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

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
func (u *User) Prepare() {
	hash, salt := hashPassword(u.Password)
	u.ID = 0
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.PasswordHash = hash
	u.Salt = salt
	u.Disabled = false
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	fmt.Println("erro:", err)
	return err
}
