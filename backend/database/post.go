package database

import (
	"errors"

	"github.com/alicelerias/blog-golang/models"
	"github.com/alicelerias/blog-golang/types"
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
	if cursor != "" {
		err := s.db.Debug().
			Where("posts.created_at < ?", cursor).
			Model(limit).
			Order("posts.created_at DESC").
			Limit(10).
			Find(&posts).
			Error
		if err != nil {
			return []models.Post{}, err
		}
	} else {
		err := s.db.Debug().
			Model(post).
			Order("posts.created_at DESC").
			Limit(limit).
			Find(&posts).
			Error
		if err != nil {
			return []models.Post{}, err
		}
	}
	if len(posts) > 0 {
		for i, _ := range posts {
			err := s.db.Debug().Model(&types.User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return []models.Post{}, err
			}
		}
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
	if cursor != "" {
		err := s.db.Debug().
			Where("posts.created_at > ? ", cursor).
			Order("posts.created_at DESC").
			Limit(limit).
			Where("author_id = ?", uid).
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
			Where("author_id = ?", uid).
			Find(&posts).
			Error
		if err != nil {
			return []models.Post{}, err
		}

		return posts, nil
	}
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
