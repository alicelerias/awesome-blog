package database

import (
	"fmt"

	"github.com/alicelerias/blog-golang/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetConnection(configs *config.Config) (db *gorm.DB, err error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", configs.Host, configs.Port, configs.User, configs.Database, configs.Password)
	db, err = gorm.Open(configs.Driver, DBURL)
	if err != nil {
		return
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(25)

	return
}
