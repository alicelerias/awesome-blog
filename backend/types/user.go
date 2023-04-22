package types

import "time"

type User struct {
	ID          uint32    `json:"id"`
	UserName    string    `json:"username"`
	Email       string    `json:"email"`
	Bio         string    `json:"bio"`
	Avatar      string    `json:"avatar"`
	IsFollowing bool      `json:"is_following"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
