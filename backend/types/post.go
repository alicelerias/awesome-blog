package types

import "time"

type Post struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Img        string    `json:"img"`
	Author     User      `json:"author"`
	AuthorID   uint32    `json:"author_id"`
	IsFavorite bool      `json:"is_favorite"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
