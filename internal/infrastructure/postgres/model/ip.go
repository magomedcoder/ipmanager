package model

import (
	"gorm.io/gorm"
)

type Ip struct {
	gorm.Model
	SubnetID    uint   `gorm:"column:subnet_id"`
	Ip          string `gorm:"column:ip"`
	CustomerID  *uint  `gorm:"column:customer_id;DEFAULT:NULL"`
	Description string `gorm:"column:description;DEFAULT:NULL"`

	Subnet   Subnet    `gorm:"foreignKey:SubnetID"`
	Customer *Customer `gorm:"foreignKey:CustomerID"`
}

func (Ip) TableName() string {
	return "ip_addresses"
}
