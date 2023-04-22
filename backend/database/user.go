package database

import (
	"errors"

	"github.com/alicelerias/blog-golang/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

func (s *PostgresDBRepository) CreateUser(ctx context.Context, user *models.User) error {
	err := s.db.Debug().Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresDBRepository) FindAllUsers(ctx context.Context, user *models.User) (*[]models.User, error) {
	var err error
	users := []models.User{}
	err = s.db.Debug().Model(&models.User{}).Order("users.created_at DESC").Limit(10).Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}
	return &users, err
}

func (s *PostgresDBRepository) FindUserByID(ctx context.Context, uid string) (user *models.User, err error) {
	user = &models.User{}
	err = s.db.First(user, uid).Error
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("User Not Found")
	}
	return
}

func (s *PostgresDBRepository) GetUser(ctx context.Context, username string) (user *models.User, err error) {
	user = &models.User{}
	err = s.db.First(user, "user_name = ?", username).Error
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("User Not Found")
	}
	return
}

func (s *PostgresDBRepository) UpdateUser(ctx context.Context, values interface{}, uid string) (u *models.User, err error) {
	u = &models.User{}
	err = s.db.Table("users").
		Where("id = ?", uid).
		UpdateColumns(values).
		Take(u).
		Error
	return
}

func (s *PostgresDBRepository) DeleteUser(ctx context.Context, uid string) (err error) {
	user := &models.User{}
	err = s.db.Delete(user, uid).Error
	if err != nil {
		return errors.New("Error")
	}
	return nil
}
