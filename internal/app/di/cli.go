package di

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/internal/cli"
	"github.com/magomedcoder/ipmanager/internal/cli/handler"
)

var CIProviderSet = wire.NewSet(
	wire.Struct(new(cli.AppProvider), "*"),
	wire.Struct(new(handler.Migrate), "*"),
)
