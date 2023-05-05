package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Feed(cursor string, followerId string) ([]models.Post, error) {
	posts := []models.Post{}
	limit := s.GetLimit()
	if cursor != "" {
		err := s.db.Debug().
			Joins("JOIN users ON posts.author_id = users.id JOIN followings ON users.id = followings.following_id").
			Where("posts.created_at < ? ", cursor).
			Where("followings.follower_id = ? OR posts.author_id = ?", followerId, followerId).
			Order("posts.created_at DESC").
			Limit(limit).
			Find(&posts).
			Error
		if err != nil {
			return []models.Post{}, err
		}

		return posts, nil
	} else {
		err := s.db.Debug().
			Order("posts.created_at DESC").
			Limit(limit).
			Joins("JOIN users ON posts.author_id = users.id JOIN followings ON users.id = followings.following_id").
			Where("followings.follower_id = ?  OR posts.author_id = ?", followerId, followerId).
			Find(&posts).
			Error
		if err != nil {
			return []models.Post{}, err
		}

		return posts, nil
	}
}
