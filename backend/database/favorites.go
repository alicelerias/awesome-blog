package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Favorite(favorite *models.Favorite) error {
	err := s.db.Debug().Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) Unfavorite(postId uint32, userId uint32) error {
	favorite := *&models.Favorite{}
	err := s.db.Debug().Where("post_id = ? AND user_id =?", postId, userId).Delete(favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) GetFavorite(postId string, userId string) bool {
	favorite := *&models.Favorite{}
	err := s.db.Debug().Where("post_id = ? AND user_id = ?", postId, userId).Find(&favorite).Error
	if err != nil {
		return false
	}

	return true
}

func (s *PostgresDBRepository) GetFavoritesByPost(postId uint32) (*[]models.Favorite, error) {
	favorites := []models.Favorite{}
	limit := s.GetLimit()
	err := s.db.Debug().Model(&models.Favorite{}).Limit(limit).Where("post_id = ?", postId).Find(&favorites).Error
	if err != nil {
		return &[]models.Favorite{}, err
	}
	return &favorites, nil
}

func (s *PostgresDBRepository) GetFavoritesPostsByUser(cursor string, userId uint32) ([]models.Post, error) {
	posts := []models.Post{}
	limit := s.GetLimit()
	if cursor != "" {
		err := s.db.Debug().
			Where("posts.created_at > ?", cursor).
			Limit(limit).
			Order("favorites.created_at desc").
			Joins("JOIN favorites ON favorites.post_id = posts.id").
			Where("favorites.user_id = ?", userId).
			Find(&posts).
			Error

		if err != nil {
			return []models.Post{}, err
		}
	} else {
		err := s.db.Debug().
			Limit(limit).
			Order("favorites.created_at desc").
			Joins("JOIN favorites ON favorites.post_id = posts.id").
			Where("favorites.user_id = ?", userId).
			Find(&posts).
			Error

		if err != nil {
			return []models.Post{}, err
		}
	}

	return posts, nil
}
