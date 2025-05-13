//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/internal/cli"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/infrastructure"
	"github.com/magomedcoder/ipmanager/internal/provider"
	"github.com/magomedcoder/ipmanager/internal/usecase"
)

var ProviderSet = wire.NewSet(
	provider.NewPostgresClient,
	infrastructure.ProviderSet,
	usecase.ProviderSet,
)

func NewGrpcInjector(conf *config.Config) *AppProvider {
	panic(wire.Build(ProviderSet, GRPCProviderSet))
}

func NewCliInjector(conf *config.Config) *cli.AppProvider {
	panic(wire.Build(ProviderSet, CIProviderSet))
}
