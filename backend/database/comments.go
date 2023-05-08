package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) CreateComment(comment *models.Comment) error {
	err := s.db.Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) DeleteComment(id uint32, authorId uint32) error {
	comment := models.Comment{}
	err := s.db.Where("id = ? AND author_id = ?", id, authorId).Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) GetPostComments(cursor string, postId uint32) ([]models.Comment, error) {
	comments := []models.Comment{}
	limit := s.GetLimit()
	query := s.db.
		Preload("Author").
		Order("comments.created_at DESC").
		Limit(limit).
		Where("post_id = ?", postId)

	if cursor != "" {
		query = query.Where("comments.created_at > ?", cursor)
	}

	err := query.Find(&comments).Error
	if err != nil {
		return []models.Comment{}, err
	}

	return comments, nil
}
