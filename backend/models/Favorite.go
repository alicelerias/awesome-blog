package models

import "time"

type Favorite struct {
	UserId    uint32    `gorm:"primary_key;auto_increment:false" json:"user_id"`
	User      User      `json:"user"`
	PostId    uint32    `gorm:"primary_key;auto_incremente:false" json:"post_id"`
	Post      Post      `json:"post"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
