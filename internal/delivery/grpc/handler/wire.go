package handler

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/api/pb"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(pb.UnimplementedUserServiceServer), "*"),

	wire.Struct(new(UserHandler), "*"),

	wire.Bind(new(pb.UserServiceServer), new(*UserHandler)),
)
