package database

import (
	"context"

	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Follow(ctx context.Context, following *models.Following) error {
	err := s.db.Debug().Create(&following).Error
	if err != nil {
		return err
	}

	if following.FollowerID != 0 {
		err = s.db.Debug().Model(&models.User{}).Where("id = ?", following.FollowerID).Take(&following.Follower).Error
		if err != nil {
			return err
		}
	}

	if following.FollowingID != 0 {
		err = s.db.Debug().Model(&models.User{}).Where("id = ?", following.FollowingID).Take(&following.Following).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *PostgresDBRepository) GetFollows(ctx context.Context, following *models.Following) (*[]models.Following, error) {
	followings := []models.Following{}
	err := s.db.Debug().Model(&following).Limit(100).Find(&followings).Error
	if err != nil {
		return &[]models.Following{}, err
	}

	return &followings, nil
}

func (s *PostgresDBRepository) Unfollow(ctx context.Context, followerId string, followingId string) error {
	following := models.Following{}
	err := s.db.Debug().Where("follower_id = ? AND following_id = ?", followerId, followingId).Delete(&following).Error
	if err != nil {
		return err
	}
	return nil
}
