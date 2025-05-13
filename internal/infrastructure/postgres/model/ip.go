package model

import (
	"gorm.io/gorm"
)

type Ip struct {
	gorm.Model
	Ip          string `gorm:"column:ip"`
	VlanID      uint   `gorm:"column:vlan_id"`
	CustomerID  uint   `gorm:"column:customer_id"`
	Description string `gorm:"column:description"`

	Vlan     Vlan     `gorm:"foreignKey:VlanID"`
	Customer Customer `gorm:"foreignKey:CustomerID"`
}

func (Ip) TableName() string {
	return "ips"
}
