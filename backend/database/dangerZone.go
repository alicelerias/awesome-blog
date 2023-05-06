package database

import (
	"log"

	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) DeleteUsersTable() error {

	if err := s.db.Delete(&models.User{}).Error; err != nil {
		log.Fatalf("Failed to clean up users table: %v", err)
	}

	return nil
}

func (s *PostgresDBRepository) DeletePostsTable() error {
	if err := s.db.Delete(&models.Post{}).Error; err != nil {
		log.Fatalf("Failed to clean up posts table: %v", err)
	}
	return nil
}

func (s *PostgresDBRepository) DeleteFollowingsTable() error {
	if err := s.db.Delete(&models.Following{}).Error; err != nil {
		log.Fatalf("Failed to clean up followings table: %v", err)
	}
	return nil
}

func (s *PostgresDBRepository) DeleteCommentsTable() error {
	if err := s.db.Delete(&models.Comment{}).Error; err != nil {
		log.Fatalf("Failed to clean up comments table: %v", err)
	}
	return nil
}

func (s *PostgresDBRepository) DeleteFavoritesTable() error {
	if err := s.db.Delete(&models.Favorite{}).Error; err != nil {
		log.Fatalf("Failed to clean up favorites table: %v", err)
	}
	return nil
}
