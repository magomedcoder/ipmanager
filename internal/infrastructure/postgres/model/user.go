package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;NOT NULL"`
	Password string `gorm:"column:password;NOT NULL"`
	Name     string `gorm:"column:name;NOT NULL"`
	Surname  string `gorm:"column:surname;NOT NULL"`
}

func (User) TableName() string {
	return "users"
}
