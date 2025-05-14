package model

type ServiceVlan struct {
	ServiceID uint `gorm:"primaryKey"`
	VlanID    uint `gorm:"primaryKey"`

	Service Service `gorm:"foreignKey:ServiceID"`
	Vlan    Vlan    `gorm:"foreignKey:VlanID"`
}

func (ServiceVlan) TableName() string {
	return "service_vlans"
}
