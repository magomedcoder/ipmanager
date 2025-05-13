package infrastructure

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres"
)

var ProviderSet = wire.NewSet(
	postgres.ProviderSet,
)
