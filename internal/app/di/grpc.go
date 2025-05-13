package di

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/handler"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/middleware"
	"github.com/magomedcoder/ipmanager/internal/server"
)

var GRPCProviderSet = wire.NewSet(
	wire.Struct(new(server.AppProvider), "*"),
	handler.ProviderSet,
	middleware.ProviderSet,
)
