package di

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/api/pb"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/handler"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/middleware"
)

type AppProvider struct {
	Middleware  middleware.AuthMiddleware
	UserHandler pb.UserServiceServer
	IpHandler   pb.IpServiceServer
}

var GRPCProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	handler.ProviderSet,
	middleware.ProviderSet,
)
