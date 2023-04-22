package database

import (
	"context"

	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Feed(ctx context.Context, cursor string, followerId string) ([]models.Post, error) {
	posts := []models.Post{}
	if cursor != "" {
		err := s.db.Debug().
			Where("posts.created_at > ? ", cursor).
			Order("posts.created_at ASC").
			Limit(10).
			Joins("JOIN users ON posts.author_id = users.id JOIN followings ON users.id = followings.following_id").
			Where("followings.follower_id = ?", followerId).
			Find(&posts).
			Error
		if err != nil {
			return []models.Post{}, err
		}

		return posts, nil
	} else {
		err := s.db.Debug().
			Order("posts.created_at ASC").
			Limit(10).
			Joins("JOIN users ON posts.author_id = users.id JOIN followings ON users.id = followings.following_id").
			Where("followings.follower_id = ?", followerId).
			Find(&posts).
			Error
		if err != nil {
			return []models.Post{}, err
		}

		return posts, nil
	}
}
