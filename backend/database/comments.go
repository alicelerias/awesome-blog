package database

import (
	"github.com/alicelerias/blog-golang/models"
	"github.com/alicelerias/blog-golang/types"
)

func (s *PostgresDBRepository) CreateComment(comment *models.Comment) error {
	err := s.db.Debug().Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) DeleteComment(id uint32, authorId uint32) error {
	comment := models.Comment{}
	err := s.db.Debug().Where("id = ? AND author_id = ?", id, authorId).Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) GetPostComments(cursor string, postId uint32) ([]models.Comment, error) {
	comments := []models.Comment{}
	limit := s.GetLimit()
	if cursor != "" {
		err := s.db.Debug().
			Where("comments.created_at > ?", cursor).
			Order("comments.created_at DESC").
			Limit(limit).
			Where("post_id = ?", postId).
			Find(&comments).
			Error
		if err != nil {
			return []models.Comment{}, err
		}
	} else {
		err := s.db.Debug().
			Order("comments.created_at DESC").
			Limit(limit).Where("post_id = ?", postId).
			Find(&comments).
			Error
		if err != nil {
			return []models.Comment{}, err
		}
	}
	if len(comments) > 0 {
		for i, _ := range comments {
			err := s.db.Debug().Model(&types.User{}).Where("id = ?", comments[i].AuthorId).Take(&comments[i].Author).Error
			if err != nil {
				return []models.Comment{}, err
			}
		}
	}
	return comments, nil
}
