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
		&model.Ip{},
	); err != nil {
		return err
	}

	return nil
}

func (m *Migrate) TestData(ctx context.Context) error {
	m.Db.Create(&model.User{
		Username: "admin",
		Password: "$2a$10$BWRbkV5MMg4JBPkeD30wO.qv1.HJcZfRZbY2uCCJnEc4YD1X7xI2i",
		Name:     "Admin",
		Surname:  "Admin",
	})

	return nil
}
