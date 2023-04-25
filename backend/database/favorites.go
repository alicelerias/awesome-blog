package database

import (
	"context"

	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Favorite(ctx context.Context, favorite *models.Favorite) error {
	err := s.db.Debug().Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) Unfavorite(ctx context.Context, postId uint32, userId uint32) error {
	favorite := *&models.Favorite{}
	err := s.db.Debug().Where("post_id = ? AND user_id =?", postId, userId).Delete(favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) GetFavorite(ctx context.Context, postId string, userId string) bool {
	favorite := *&models.Favorite{}
	err := s.db.Debug().Where("post_id = ? AND user_id = ?", postId, userId).Find(&favorite).Begin().Error
	if err != nil {
		return false
	}

	return true
}

func (s *PostgresDBRepository) GetFavoritesByPost(ctx context.Context, postId uint32) (*[]models.Favorite, error) {
	favorites := []models.Favorite{}
	err := s.db.Debug().Model(&models.Favorite{}).Limit(100).Where("post_id = ?", postId).Find(&favorites).Error
	if err != nil {
		return &[]models.Favorite{}, err
	}
	return &favorites, nil
}

func (s *PostgresDBRepository) GetFavoritesPostsByUser(ctx context.Context, cursor string, userId uint32) ([]models.Post, error) {
	posts := []models.Post{}
	if cursor != "" {
		err := s.db.Debug().
			Where("posts.created_at > ?", cursor).
			Limit(10).
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
			Limit(10).
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
