package entity

type Service struct {
	Id    int64
	Name  string
	Ips   []ServiceIp
	Vlans []ServiceVlan
}

type ServiceIp struct {
	Id int64
	Ip string
}

type ServiceVlan struct {
	Id   int64
	Name string
}
