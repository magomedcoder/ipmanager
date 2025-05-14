package model

import (
	"gorm.io/gorm"
)

type Subnet struct {
	gorm.Model
	Ip          string `gorm:"column:ip"`
	VlanID      *int64 `gorm:"column:vlan_id;DEFAULT:NULL"`
	Description string `gorm:"column:description"`

	Vlan *Vlan `gorm:"foreignKey:VlanID"`
}

func (Subnet) TableName() string {
	return "subnets"
}
