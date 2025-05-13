//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/infrastructure"
	"github.com/magomedcoder/ipmanager/internal/provider"
	"github.com/magomedcoder/ipmanager/internal/server"
	"github.com/magomedcoder/ipmanager/internal/usecase"
)

var ProviderSet = wire.NewSet(
	provider.NewPostgresClient,
	infrastructure.ProviderSet,
	usecase.ProviderSet,
)

func NewGrpcInjector(conf *config.Config) *server.AppProvider {
	panic(wire.Build(ProviderSet, GRPCProviderSet))
}
