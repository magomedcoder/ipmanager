package model

import "gorm.io/gorm"

type Vlan struct {
	gorm.Model
	Name string `gorm:"column:name;NOT NULL"`

	Services []Service `gorm:"many2many:service_vlans"`
}

func (Vlan) TableName() string {
	return "vlans"
}
