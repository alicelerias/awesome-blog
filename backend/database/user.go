package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) CreateUser(user *models.User) error {
	err := s.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresDBRepository) FindAllUsers(user *models.User) (*[]models.User, error) {
	limit := s.GetLimit()
	users := []models.User{}
	err := s.db.Model(&models.User{}).Order("users.created_at DESC").Limit(limit).Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}
	return &users, err
}

func (s *PostgresDBRepository) FindUserByID(uid string) (*models.User, error) {
	user := &models.User{}
	if err := s.db.First(user, uid).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (s *PostgresDBRepository) GetUser(username string) (*models.User, error) {
	user := &models.User{}
	if err := s.db.First(user, "user_name = ?", username).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (s *PostgresDBRepository) UpdateUser(values interface{}, uid string) (*models.User, error) {
	user := &models.User{}
	if err := s.db.Table("users").
		Where("id = ?", uid).
		UpdateColumns(values).
		Take(user).
		Error; err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (s *PostgresDBRepository) DeleteUser(uid string) error {
	user := &models.User{}
	if err := s.db.Delete(user, uid).Error; err != nil {
		return err
	}
	return nil
}
