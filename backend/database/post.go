package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) CreatePost(post *models.Post) error {
	err := s.db.Preload("Author").Create(&post).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) GetPosts(cursor string, post *models.Post) ([]models.Post, error) {
	posts := []models.Post{}
	limit := s.GetLimit()
	query := s.db.Preload("Author").Joins("JOIN popularity_score on posts.author_id = popularity_score.id").Order("popularity_score.score DESC, posts.created_at DESC").
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

func (s *PostgresDBRepository) GetPost(id string) (*models.Post, error) {
	post := &models.Post{}
	if err := s.db.Preload("Author").First(post, id).Error; err != nil {
		return &models.Post{}, err
	}

	return post, nil

}

func (s *PostgresDBRepository) GetPostsByUser(post *models.Post, cursor string, uid string) ([]models.Post, error) {
	posts := []models.Post{}
	limit := s.GetLimit()
	query := s.db.Preload("Author").Where("author_id = ?", uid).Order("posts.created_at DESC").
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

func (s *PostgresDBRepository) UpdatePost(values interface{}, id string) (*models.Post, error) {
	post := &models.Post{}
	if err := s.db.Preload("Author").Table("posts").Where("id = ?", id).UpdateColumns(values).Take(post).Error; err != nil {
		return &models.Post{}, err
	}

	return post, nil

}
func (s *PostgresDBRepository) DeletePost(id string) error {
	post := &models.Post{}
	err := s.db.Delete(post, id).Error
	if err != nil {
		return err
	}
	return nil
}
