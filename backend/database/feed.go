package database

import (
	"context"

	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Feed(ctx context.Context, followerId string) (*[]models.Post, error) {
	posts := []models.Post{}
	err := s.db.Debug().
		Limit(1).
		Order("posts.created_at desc").
		Joins("JOIN users ON posts.author_id = users.id JOIN followings ON users.id = followings.following_id").
		Where("followings.follower_id = ?", followerId).
		Find(&posts).
		Error
	if err != nil {
		return &[]models.Post{}, err
	}

	return &posts, nil
}
