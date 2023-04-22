package database

import (
	"github.com/alicelerias/blog-golang/models"
	"github.com/jinzhu/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Following{},
		&models.Favorite{},
		&models.Comment{},
	)
}
