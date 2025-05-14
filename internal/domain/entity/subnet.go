package entity

type Subnet struct {
	Id           int64
	Ip           string
	VlanId       int64
	VlanName     string
	CustomerId   int64
	CustomerName string
	Description  string
}
