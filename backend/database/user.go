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
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName       string    `gorm:"size:255;not null;unique" json:"username"`
	Email          string    `gorm:"size:100;not null;unique" json:"email"`
	Password       string    `gorm:"size:100;not null;" json:"password"`
	PasswordHash   []byte    `gorm:"not null" json:"pwd_hash"`
	Salt           []byte    `gorm:"not null" json:"salt"`
	Disabled       bool      `gorm:"not null" json:"disabled"`
	DisabledReason string    `gorm:"not null" json:"disabled_reason"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// func Hash(password string) ([]byte, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }

// func VerifyPassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// func (u *User) BeforeSave() error {
// 	hashedPassword, err := Hash(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }

func (s *PostgresDBRepository) SaveUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := s.db.Debug().Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (s *PostgresDBRepository) FindAllUsers(ctx context.Context, user *models.User) (*[]models.User, error) {
	var err error
	users := []models.User{}
	err = s.db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
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

func (s *PostgresDBRepository) FindUserByID(ctx context.Context, uid int64) (user *models.User, err error) {
	fmt.Println(user) // nil
	user = &models.User{}
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
