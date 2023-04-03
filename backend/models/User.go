package models

import (
	"time"
)

type User struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	UserName string `gorm:"size:255;not null;unique" json:"username" validate:"required,min=4,max=15"`
	Email    string `gorm:"size:100;not null;unique" json:"email" validate:"required,email"`
	// Bio string `...`
	Password       string    `gorm:"size:100;not null;" json:"password" validate:"required,passwd,min=6,max=15"`
	PasswordHash   []byte    `gorm:"not null" json:"pwd_hash"`
	Salt           []byte    `gorm:"not null" json:"salt"`
	Disabled       bool      `gorm:"not null" json:"disabled"`
	DisabledReason string    `json:"disabled_reason"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
