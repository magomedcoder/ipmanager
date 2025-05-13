package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/domain/entity"
	postgresModel "github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	postgresRepo "github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type IVlanUseCase interface {
	Create(ctx context.Context, opt *VlanOpt) (*entity.Vlan, error)

	GetVlans(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Vlan, error)

	GetById(ctx context.Context, id int64) (*entity.Vlan, error)
}

var _ IVlanUseCase = (*VlanUseCase)(nil)

type VlanUseCase struct {
	Conf     *config.Config
	VlanRepo postgresRepo.IVlanRepository
}

type VlanOpt struct {
	Id   int64
	Name string
}

func (v *VlanUseCase) Create(ctx context.Context, opt *VlanOpt) (*entity.Vlan, error) {
	vlan, err := v.VlanRepo.Create(ctx, &postgresModel.Vlan{
		Name: opt.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать Vlan")
	}

	return &entity.Vlan{
		Id:   int64(vlan.ID),
		Name: vlan.Name,
	}, nil
}

func (v *VlanUseCase) GetVlans(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Vlan, error) {
	vlans, err := v.VlanRepo.GetVlans(ctx, arg...)
	if err != nil {
		return nil, err
	}

	items := make([]*entity.Vlan, 0)
	for _, item := range vlans {
		items = append(items, &entity.Vlan{
			Id:   int64(item.ID),
			Name: item.Name,
		})
	}

	return items, nil
}

func (v *VlanUseCase) GetById(ctx context.Context, id int64) (*entity.Vlan, error) {
	vlan, err := v.VlanRepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Не удалось получить vlan: %v", id))
	}

	return &entity.Vlan{
		Id:   int64(vlan.ID),
		Name: vlan.Name,
	}, nil
}
