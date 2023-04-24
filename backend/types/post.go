package types

import "time"

type Post struct {
	ID             uint64    `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Img            string    `json:"img"`
	Author         User      `json:"author"`
	AuthorID       uint32    `json:"author_id"`
	IsFavorite     bool      `json:"is_favorite"`
	CommentsCount  int       `json:"comments_count"`
	FavoritesCount int       `json:"favorites_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
