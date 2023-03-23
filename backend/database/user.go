package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/alicelerias/blog-golang/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *PostgresDBRepository) SaveUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := s.db.Debug().Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (s *PostgresDBRepository) FindAllUsers(ctx context.Context, user *models.User) (*[]User, error) {
	var err error
	users := []User{}
	err = s.db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// func typeExample() {
// 	var user User
// 	fmt.Println(user) // nil
// 	// var userB User = nil
// 	// fmt.Println(user) // nil
// 	var userA User = User{}
// 	fmt.Println(userA) // User{}
// 	userB := User{}
// 	fmt.Println(userB) // User{}

// 	var userC *User = &User{}
// 	fmt.Println(userC) // User{}

// 	userD := &User{}
// 	fmt.Println(userD) // User{}
// 	GetValue(userD)

// }

func (s *PostgresDBRepository) FindUserByID(ctx context.Context, uid int64) (user *User, err error) {
	fmt.Println(user) // nil
	user = &User{}
	err = s.db.First(user, uid).Error
	if gorm.IsRecordNotFoundError(err) {
		err = errors.New("User Not Found")
	}
	return
}

func (s *PostgresDBRepository) DeleteAUser(ctx context.Context, uid int64) (err error) {
	user := &models.User{}
	err = s.db.Delete(user, uid).Error
	if err != nil {
		return errors.New("Error")
	}
	return nil
}
