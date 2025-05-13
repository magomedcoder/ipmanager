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

	wire.Bind(new(repository.ICustomerRepository), new(*repository.CustomerRepository)),
	repository.NewCustomerRepository,

	wire.Bind(new(repository.IVlanRepository), new(*repository.VlanRepository)),
	repository.NewVlanRepository,

	wire.Bind(new(repository.IServiceRepository), new(*repository.ServiceRepository)),
	repository.NewServiceRepository,
)
