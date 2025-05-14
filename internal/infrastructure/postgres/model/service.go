package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name string `gorm:"column:name;NOT NULL"`

	Ips   []Ip   `gorm:"foreignKey:ServiceID"`
	Vlans []Vlan `gorm:"many2many:service_vlans"`
}

func (Service) TableName() string {
	return "services"
}
