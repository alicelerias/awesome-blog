package database

import (
	"context"

	"github.com/alicelerias/blog-golang/models"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetHome() (err error)
	// VerifyPassword() (err error)
	// BeforeSave() (err error)
	SaveUser(context.Context, *models.User) (*models.User, error)
	FindAllUsers(context.Context, *models.User) (*[]User, error)
	FindUserByID(context.Context, int64) (*User, error)
	DeleteAUser(context.Context, int64) error
}

type PostgresDBRepository struct {
	db *gorm.DB
}

var postgresDBRepository Repository = &PostgresDBRepository{}

func NewPostgresDBRepository(db *gorm.DB) *PostgresDBRepository {
	return &PostgresDBRepository{db: db}
}

// connection := GetConnection()
// 	db, _, _:= sqlmock.New()
// 	defer db.Close()
// 	mockDB, err := gorm.Open("postgres", db)
