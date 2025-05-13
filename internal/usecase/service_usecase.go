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

type IServiceUseCase interface {
	Create(ctx context.Context, opt *ServiceOpt) (*entity.Service, error)

	GetServices(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Service, error)

	GetById(ctx context.Context, id int64) (*entity.Service, error)
}

var _ IServiceUseCase = (*ServiceUseCase)(nil)

type ServiceUseCase struct {
	Conf        *config.Config
	ServiceRepo postgresRepo.IServiceRepository
}

type ServiceOpt struct {
	Id   int64
	Name string
}

func (v *ServiceUseCase) Create(ctx context.Context, opt *ServiceOpt) (*entity.Service, error) {
	service, err := v.ServiceRepo.Create(ctx, &postgresModel.Service{
		Name: opt.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать Service")
	}

	return &entity.Service{
		Id:   int64(service.ID),
		Name: service.Name,
	}, nil
}

func (v *ServiceUseCase) GetServices(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Service, error) {
	services, err := v.ServiceRepo.GetServices(ctx, arg...)
	if err != nil {
		return nil, err
	}

	items := make([]*entity.Service, 0)
	for _, item := range services {
		items = append(items, &entity.Service{
			Id:   int64(item.ID),
			Name: item.Name,
		})
	}

	return items, nil
}

func (v *ServiceUseCase) GetById(ctx context.Context, id int64) (*entity.Service, error) {
	service, err := v.ServiceRepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Не удалось получить service: %v", id))
	}

	return &entity.Service{
		Id:   int64(service.ID),
		Name: service.Name,
	}, nil
}
