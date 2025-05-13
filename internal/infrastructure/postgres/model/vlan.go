package model

import "gorm.io/gorm"

type Vlan struct {
	gorm.Model
	Name string `gorm:"column:name;NOT NULL"`
}

func (Vlan) TableName() string {
	return "vlans"
}
