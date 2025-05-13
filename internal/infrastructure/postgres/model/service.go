package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name string `gorm:"column:name;NOT NULL"`
}

func (Service) TableName() string {
	return "services"
}
