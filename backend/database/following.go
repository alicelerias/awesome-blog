package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Follow(following *models.Following) error {
	err := s.db.Create(&following).Error
	if err != nil {
		return err
	}

	if following.FollowerID != 0 {
		err = s.db.Model(&models.User{}).Where("id = ?", following.FollowerID).Take(&following.Follower).Error
		if err != nil {
			return err
		}
	}

	if following.FollowingID != 0 {
		err = s.db.Model(&models.User{}).Where("id = ?", following.FollowingID).Take(&following.Following).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *PostgresDBRepository) GetFollows(following *models.Following) (*[]models.Following, error) {
	followings := []models.Following{}
	limit := s.GetLimit()
	err := s.db.Model(&following).Limit(limit).Find(&followings).Error
	if err != nil {
		return &[]models.Following{}, err
	}

	return &followings, nil
}

func (s *PostgresDBRepository) IsFollowing(followerId string, followingId string) bool {
	following := models.Following{}
	err := s.db.Where("follower_id = ? AND following_id = ?", followerId, followingId).Find(&following).Error
	if err != nil {
		return false
	}
	return true
}

func (s *PostgresDBRepository) Unfollow(followerId string, followingId string) error {
	following := models.Following{}
	err := s.db.Where("follower_id = ? AND following_id = ?", followerId, followingId).Delete(&following).Error
	if err != nil {
		return err
	}
	return nil
}
