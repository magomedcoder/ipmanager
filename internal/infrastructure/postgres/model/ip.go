package model

import (
	"gorm.io/gorm"
)

type Ip struct {
	gorm.Model
	Ip string `gorm:"column:ip"`
}

func (Ip) TableName() string {
	return "ips"
}
