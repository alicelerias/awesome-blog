package models

type Following struct {
	FollowerID  uint32 `gorm:"primary_key;auto_increment:false"`
	Follower    User
	FollowingID uint32 `gorm:"primary_key;auto_increment:false"`
	Following   User
}
