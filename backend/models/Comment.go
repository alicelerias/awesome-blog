package models

import "time"

type Comment struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	PostId    uint32    `json:"post_id"`
	Post      Post      `json:"post"`
	AuthorId  uint32    `gorm:"not null" json:"author_id"`
	Author    User      `json:"author"`
	Content   string    `gorm:"size:300;not null" json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
