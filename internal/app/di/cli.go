package di

import (
	"github.com/magomedcoder/ipmanager/internal/cli/handler"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/repository"
	"github.com/magomedcoder/ipmanager/internal/provider"
	"github.com/magomedcoder/ipmanager/internal/usecase"
)

type CliProvider struct {
	Conf    *config.Config
	Migrate *handler.Migrate
}

func NewCliInjector(conf *config.Config) *CliProvider {
	db := provider.NewPostgresClient(conf)
	userRepository := repository.NewUserRepository(db)
	userSessionRepository := repository.NewUserSessionRepository(db)
	userUseCase := &usecase.UserUseCase{
		Conf:            conf,
		UserRepo:        userRepository,
		UserSessionRepo: userSessionRepository,
	}
	migrate := &handler.Migrate{
		Conf:        conf,
		Db:          db,
		UserUseCase: userUseCase,
	}
	appProvider := &CliProvider{
		Conf:    conf,
		Migrate: migrate,
	}

	return appProvider
}
