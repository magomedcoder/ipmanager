package model

import (
	"gorm.io/gorm"
)

type UserSession struct {
	gorm.Model
	AccessToken string `gorm:"column:access_token;NOT NULL"`
	ExpiresAt   int64  `gorm:"column:expires_at;NOT NULL"`
	UserId      uint   `gorm:"column:user_id;NOT NULL"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}
