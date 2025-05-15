package repository

import (
	"context"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"gorm.io/gorm"
	"log"
)

type IServiceRepository interface {
	Create(ctx context.Context, service *model.Service) (*model.Service, error)

	GetServices(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Service, error)

	GetById(ctx context.Context, id int64) (*model.Service, error)
}

var _ IServiceRepository = (*ServiceRepository)(nil)

type ServiceRepository struct {
	gormutil.Repo[model.Service]
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{Repo: gormutil.NewRepo[model.Service](db)}
}

func (s *ServiceRepository) Create(ctx context.Context, service *model.Service) (*model.Service, error) {
	if err := s.Repo.Create(ctx, service); err != nil {
		log.Printf("Не удалось создать: %s", err)
		return nil, err
	}

	return service, nil
}

func (s *ServiceRepository) GetServices(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Service, error) {
	services, err := s.FindAll(ctx, arg...)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s *ServiceRepository) GetById(ctx context.Context, id int64) (*model.Service, error) {
	service, err := s.Repo.FindByWhereWithQuery(ctx, "id = ?", []any{id}, func(db *gorm.DB) {
		db.Preload("Ips")
	})
	if err != nil {
		return nil, err
	}

	return service, nil
}
