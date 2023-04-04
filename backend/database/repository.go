package database

import (
	"context"

	"github.com/alicelerias/blog-golang/models"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetHome() (err error)
	CreateUser(context.Context, *models.User) error
	FindAllUsers(context.Context, *models.User) (*[]models.User, error)
	FindUserByID(context.Context, string) (*models.User, error)
	GetUser(context.Context, string) (*models.User, error)
	UpdateUser(context.Context, interface{}, string) (*models.User, error)
	DeleteUser(context.Context, string) error

	CreatePost(context.Context, *models.Post) error
	GetPosts(context.Context, *models.Post) (*[]models.Post, error)
	GetPost(context.Context, string) (*models.Post, error)
	UpdatePost(context.Context, interface{}, string) (*models.Post, error)
	DeletePost(context.Context, string) error

	Follow(context.Context, *models.Following) error
	GetFollows(context.Context, *models.Following) (*[]models.Following, error)
	Unfollow(context.Context, string, string) error
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
