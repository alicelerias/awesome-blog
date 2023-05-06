package database

import (
	"errors"

	"github.com/alicelerias/blog-golang/models"
	"github.com/jinzhu/gorm"
)

func (s *PostgresDBRepository) CreatePost(post *models.Post) error {
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

func (s *PostgresDBRepository) GetPosts(cursor string, post *models.Post) ([]models.Post, error) {
	posts := []models.Post{}
	limit := s.GetLimit()
	query := s.db.Preload("Author").Order("posts.created_at DESC").
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

func (s *PostgresDBRepository) GetPost(id string) (post *models.Post, err error) {
	post = &models.Post{}
	err = s.db.First(post, id).Error
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("Post Not Found")
	}
	if post.ID != 0 {
		err = s.db.Debug().Model(&models.Post{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
		if err != nil {
			return &models.Post{}, err
		}
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

func (s *PostgresDBRepository) UpdatePost(values interface{}, id string) (post *models.Post, err error) {
	post = &models.Post{}
	err = s.db.Table("posts").Where("id = ?", id).UpdateColumns(values).Take(post).Error
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("Post Not Found")
	}
	if post.ID != 0 {
		err = s.db.Debug().Model(&models.Post{}).Where("id = ?", post.AuthorID).Take(&post.Author).Error
		if err != nil {
			return
		}
	}
	return

}
func (s *PostgresDBRepository) DeletePost(id string) error {
	post := &models.Post{}
	err := s.db.Debug().Delete(post, id).Error
	if err != nil {
		return err
	}
	return nil
}
