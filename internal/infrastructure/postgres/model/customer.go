package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `gorm:"column:name;NOT NULL"`
	Surname string `gorm:"column:surname;NOT NULL"`
}

func (Customer) TableName() string {
	return "customers"
}
