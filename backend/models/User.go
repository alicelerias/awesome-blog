package models

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:255;not null;unique" json:"username" validate:"required,min=4,max=15"`
	Email     string    `gorm:"size:100;not null;unique" json:"email" validate:"required,email"`
	Password  string    `gorm:"size:100;not null;" json:"password" validate:"required,min=6,max=15"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) Prepare() {
	u.ID = 0
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	fmt.Println("erro:", err)
	return err
}
