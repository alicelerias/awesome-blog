package database

import (
	"errors"
	"time"

	"github.com/alicelerias/blog-golang/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (s *PostgresDBRepository) SaveUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := s.db.Debug().Create(&user).Error
	if err != nil {
		log.Error()
		return &models.User{}, err
	}

	return user, nil
}

func (s *PostgresDBRepository) FindAllUsers(ctx context.Context, user *models.User) (*[]models.User, error) {
	var err error
	users := []models.User{}
	err = s.db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		log.Error()
		return &[]models.User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAuser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"username":   u.UserName,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil

}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
