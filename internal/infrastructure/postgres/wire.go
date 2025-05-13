package postgres

import (
	"github.com/google/wire"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/repository"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(repository.IUserRepository), new(*repository.UserRepository)),
	repository.NewUserRepository,

	wire.Bind(new(repository.IUserSessionRepository), new(*repository.UserSessionRepository)),
	repository.NewUserSessionRepository,

	wire.Bind(new(repository.IIpRepository), new(*repository.IpRepository)),
	repository.NewIPRepository,
)
