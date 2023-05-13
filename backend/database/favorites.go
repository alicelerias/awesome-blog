package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Favorite(favorite *models.Favorite) error {
	err := s.db.Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) Unfavorite(postId uint32, userId uint32) error {
	favorite := *&models.Favorite{}
	err := s.db.Where("post_id = ? AND user_id =?", postId, userId).Delete(favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresDBRepository) GetFavorite(postId string, userId string) bool {
	favorite := *&models.Favorite{}
	err := s.db.Where("post_id = ? AND user_id = ?", postId, userId).Find(&favorite).Error
	if err != nil {
		return false
	}

	return true
}

func (s *PostgresDBRepository) GetFavoritesByPost(postId uint32) (*[]models.Favorite, error, int) {
	var count int
	favorites := []models.Favorite{}
	limit := s.GetLimit()
	err := s.db.Model(&models.Favorite{}).Limit(limit).Where("post_id = ?", postId).Find(&favorites).Count(&count).Error
	if err != nil {
		return &[]models.Favorite{}, err, 0
	}
	return &favorites, nil, count
}

func (s *PostgresDBRepository) GetFavoritesPostsByUser(cursor string, userId uint32) ([]models.Post, error) {
	posts := []models.Post{}
	limit := s.GetLimit()
	query := s.db.Joins("JOIN favorites ON favorites.post_id = posts.id").
		Where("favorites.user_id = ?", userId).
		Order("favorites.created_at desc").
		Limit(limit)
	if cursor != "" {
		query = query.Where("posts.created_at > ?", cursor)
	}
	err := query.Find(&posts).
		Error
	if err != nil {
		return []models.Post{}, err
	}

	return posts, nil
}
