package database

import (
	"github.com/alicelerias/blog-golang/models"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetLimit() string

	Recomendations(uid string) (*[]models.User, error)

	GetHome() (err error)
	CreateUser(*models.User) error
	FindAllUsers(*models.User) (*[]models.User, error)
	FindUserByID(string) (*models.User, error)
	GetUser(string) (*models.User, error)
	UpdateUser(interface{}, string) (*models.User, error)
	DeleteUser(string) error

	CreatePost(*models.Post) error
	GetPosts(string, *models.Post) ([]models.Post, error)
	GetPostsByUser(post *models.Post, cursor string, uid string) ([]models.Post, error)
	GetPost(string) (*models.Post, error)
	UpdatePost(interface{}, string) (*models.Post, error)
	DeletePost(string) error

	Favorite(*models.Favorite) error
	Unfavorite(postId uint32, userId uint32) error
	GetFavorite(postId string, userId string) bool
	GetFavoritesByPost(postId uint32) (*[]models.Favorite, error)
	GetFavoritesPostsByUser(cursor string, userId uint32) ([]models.Post, error)

	CreateComment(comment *models.Comment) error
	DeleteComment(id uint32, authorId uint32) error
	GetPostComments(cursor string, postId uint32) ([]models.Comment, error)

	Follow(*models.Following) error
	GetFollows(*models.Following) (*[]models.Following, error)
	IsFollowing(followerId string, followingId string) bool
	Unfollow(string, string) error

	Feed(cursor string, followerId string) ([]models.Post, error)

	// danger zone

	DeleteUsersTable() error
	DeletePostsTable() error
	DeleteFollowingsTable() error
	DeleteCommentsTable() error
	DeleteFavoritesTable() error
}

type PostgresDBRepository struct {
	db *gorm.DB
}

var postgresDBRepository Repository = &PostgresDBRepository{}

func NewPostgresDBRepository(db *gorm.DB) *PostgresDBRepository {
	return &PostgresDBRepository{db: db}
}
