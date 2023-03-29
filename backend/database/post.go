package database

import (
	"context"

	"github.com/alicelerias/blog-golang/models"
	"github.com/alicelerias/blog-golang/types"
)

func (s *PostgresDBRepository) CreatePost(ctx context.Context, post *models.Post) error {
	err := s.db.Debug().Create(&post).Error
	if err != nil {
		return err
	}

	if post.ID != 0 {
		err = s.db.Debug().Model(&models.User{}).Where("id = ?", post.ID).Take(&post.Author).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PostgresDBRepository) GetPosts(ctx context.Context, post *models.Post) (*[]models.Post, error) {
	posts := []models.Post{}
	err := s.db.Debug().Model(&models.Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]models.Post{}, err
	}
	if len(posts) > 0 {
		for i, _ := range posts {
			err := s.db.Debug().Model(&types.User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]models.Post{}, err
			}
		}
	}
	return &posts, err
}

func (s *PostgresDBRepository) DeletePost(ctx context.Context, uid string) error {
	post := &models.Post{}
	err := s.db.Debug().Delete(post, uid).Error
	if err != nil {
		return err
	}
	return nil
}
