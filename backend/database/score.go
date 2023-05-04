package database

import (
	"github.com/alicelerias/blog-golang/types"
)

func (s *PostgresDBRepository) UsersScore() ([]types.UserScore, error) {
	users := []types.UserScore{}

	err := s.db.Table("users").
		Select("users.id, users.user_name, COUNT(DISTINCT favorites.post_id) as favorites_count, COUNT(DISTINCT comments.id) as comments_count, COUNT(DISTINCT favorites.post_id) + COUNT(DISTINCT comments.id) as score").
		Joins("LEFT JOIN posts ON posts.author_id = users.id").
		Joins("LEFT JOIN favorites ON favorites.post_id = posts.id").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("users.id").
		Order("score DESC").
		Scan(&users).Error
	if err != nil {
		return []types.UserScore{}, err
	}
	return users, nil
}
