package handler

import (
	"context"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/internal/usecase"

	"gorm.io/gorm"
)

type Migrate struct {
	Conf        *config.Config
	Db          *gorm.DB
	UserUseCase usecase.IUserUseCase
}

func (m *Migrate) Migrate(ctx context.Context) error {
	if err := m.Db.AutoMigrate(
		&model.User{},
		&model.UserSession{},
		&model.Customer{},
		&model.Vlan{},
		&model.Subnet{},
		&model.Ip{},
		&model.Service{},
		&model.ServiceVlan{},
	); err != nil {
		return err
	}

	userUseCase := usecase.UserUseCase{}
	password, err := userUseCase.HashPassword("admin123")
	if err != nil {
		return err
	}

	m.Db.Create(&model.User{
		Username: "admin",
		Password: password,
		Name:     "Admin",
		Surname:  "Admin",
	})

	return nil
}
