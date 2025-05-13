package repository

import (
	"context"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"gorm.io/gorm"
	"log"
)

type IVlanRepository interface {
	Create(ctx context.Context, vlan *model.Vlan) (*model.Vlan, error)

	GetVlans(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Vlan, error)

	GetById(ctx context.Context, id int64) (*model.Vlan, error)
}

var _ IVlanRepository = (*VlanRepository)(nil)

type VlanRepository struct {
	gormutil.Repo[model.Vlan]
}

func NewVlanRepository(db *gorm.DB) *VlanRepository {
	return &VlanRepository{Repo: gormutil.NewRepo[model.Vlan](db)}
}

func (v *VlanRepository) Create(ctx context.Context, vlan *model.Vlan) (*model.Vlan, error) {
	if err := v.Repo.Create(ctx, vlan); err != nil {
		log.Printf("Не удалось создать пользователя: %s", err)
		return nil, err
	}

	return vlan, nil
}

func (v *VlanRepository) GetVlans(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Vlan, error) {
	vlans, err := v.FindAll(ctx, arg...)
	if err != nil {
		return nil, err
	}

	return vlans, nil
}

func (v *VlanRepository) GetById(ctx context.Context, id int64) (*model.Vlan, error) {
	vlan, err := v.Repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return vlan, nil
}
