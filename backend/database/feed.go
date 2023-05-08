package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Feed(cursor string, followerId string) ([]models.Post, error) {
	posts := []models.Post{}
	limit := s.GetLimit()
	query := s.db.
		Preload("Author").
		Joins("JOIN users ON posts.author_id = users.id JOIN followings ON users.id = followings.following_id").
		Where("followings.follower_id = ?", followerId).
		Order("posts.created_at DESC").
		Limit(limit)
	if cursor != "" {
		query = query.Where("posts.created_at < ?", cursor)
	}
	err := query.Find(&posts).
		Error
	if err != nil {
		return []models.Post{}, err

	}
	return posts, nil
}
