package handler

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/api/pb"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(UserHandler), "*"),
	wire.Struct(new(pb.UnimplementedUserServiceServer), "*"),
	wire.Bind(new(pb.UserServiceServer), new(*UserHandler)),

	wire.Struct(new(IpHandler), "*"),
	wire.Struct(new(pb.UnimplementedIpServiceServer), "*"),
	wire.Bind(new(pb.IpServiceServer), new(*IpHandler)),

	wire.Struct(new(CustomerHandler), "*"),
	wire.Struct(new(pb.UnimplementedCustomerServiceServer), "*"),
	wire.Bind(new(pb.CustomerServiceServer), new(*CustomerHandler)),

	wire.Struct(new(VlanHandler), "*"),
	wire.Struct(new(pb.UnimplementedVlanServiceServer), "*"),
	wire.Bind(new(pb.VlanServiceServer), new(*VlanHandler)),
)
