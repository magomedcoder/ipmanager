package entity

type Ip struct {
	Id           int64
	Ip           string
	Busy         bool
	SubnetId     int64
	SubnetName   string
	ServiceId    int64
	ServiceName  string
	CustomerId   int64
	CustomerName string
	Description  string
}
