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
	GetPosts(context.Context, string, *models.Post) ([]models.Post, error)
	GetPostsByUser(ctx context.Context, post *models.Post, cursor string, uid string) ([]models.Post, error)
	GetPost(context.Context, string) (*models.Post, error)
	UpdatePost(context.Context, interface{}, string) (*models.Post, error)
	DeletePost(context.Context, string) error

	Favorite(context.Context, *models.Favorite) error
	Unfavorite(ctx context.Context, postId uint32, userId uint32) error
	GetFavorite(ctx context.Context, postId string, userId string) bool
	GetFavoritesByPost(ctx context.Context, postId uint32) (*[]models.Favorite, error)
	GetFavoritesPostsByUser(ctx context.Context, cursor string, userId uint32) ([]models.Post, error)

	CreateComment(ctx context.Context, comment *models.Comment) error
	DeleteComment(ctx context.Context, id uint32, authorId uint32) error
	GetPostComments(ctx context.Context, postId uint32) (*[]models.Comment, error)

	Follow(context.Context, *models.Following) error
	GetFollows(context.Context, *models.Following) (*[]models.Following, error)
	IsFollowing(ctx context.Context, followerId string, followingId string) bool
	Unfollow(context.Context, string, string) error

	Feed(ctx context.Context, cursor string, followerId string) ([]models.Post, error)
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
