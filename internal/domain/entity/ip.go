package entity

type Ip struct {
	Id           int64
	Ip           string
	Busy         bool
	SubnetId     int64
	SubnetName   string
	VlanId       int64
	VlanName     string
	CustomerId   int64
	CustomerName string
	Description  string
}
