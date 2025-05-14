package repository

import (
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"gorm.io/gorm"
)

type IServiceVlanRepository interface {
}

var _ IServiceVlanRepository = (*ServiceVlanRepository)(nil)

type ServiceVlanRepository struct {
	gormutil.Repo[model.ServiceVlan]
}

func NewServiceVlanRepository(db *gorm.DB) *ServiceVlanRepository {
	return &ServiceVlanRepository{Repo: gormutil.NewRepo[model.ServiceVlan](db)}
}
