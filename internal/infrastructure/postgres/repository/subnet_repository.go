package repository

import (
	"context"
	"errors"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"gorm.io/gorm"
	"log"
)

type ISubnetRepository interface {
	Create(ctx context.Context, subnet *model.Subnet) (*model.Subnet, error)

	GetSubnets(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Subnet, error)

	GetById(ctx context.Context, id int64) (*model.Subnet, error)

	GetBySubnet(subnet string) (*model.Subnet, error)

	EditVlanById(ctx context.Context, id int64, vlanId int64) error

	EditDescriptionById(ctx context.Context, id int64, description string) error
}

var _ ISubnetRepository = (*SubnetRepository)(nil)

type SubnetRepository struct {
	gormutil.Repo[model.Subnet]
}

func NewSubnetRepository(db *gorm.DB) *SubnetRepository {
	return &SubnetRepository{Repo: gormutil.NewRepo[model.Subnet](db)}
}

func (s *SubnetRepository) Create(ctx context.Context, subnet *model.Subnet) (*model.Subnet, error) {
	if err := s.Repo.Create(ctx, subnet); err != nil {
		log.Printf("Не удалось создать subnet: %s", err)
		return nil, err
	}

	return subnet, nil
}

func (s *SubnetRepository) GetSubnets(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Subnet, error) {
	subnets, err := s.FindAll(ctx, arg...)
	if err != nil {
		return nil, err
	}

	return subnets, nil
}

func (s *SubnetRepository) GetById(ctx context.Context, id int64) (*model.Subnet, error) {
	subnet, err := s.Repo.FindByWhereWithQuery(ctx, "id = ?", []any{id}, func(db *gorm.DB) {
		db.Preload("Vlan")
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Не удалось получить: %s", err)
		}
		return nil, err
	}

	return subnet, nil
}

func (s *SubnetRepository) GetBySubnet(subnet string) (*model.Subnet, error) {
	subnetData, err := s.Repo.FindByWhere(context.TODO(), "ip = ?", subnet)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Не удалось получить пользователя: %s", err)
		}
		return nil, err
	}

	return subnetData, nil
}

func (s *SubnetRepository) EditVlanById(ctx context.Context, id int64, vlanId int64) error {
	_, err := s.Repo.UpdateById(ctx, id, map[string]any{
		"vlan_id": vlanId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *SubnetRepository) EditDescriptionById(ctx context.Context, id int64, description string) error {
	_, err := s.Repo.UpdateById(ctx, id, map[string]any{
		"description": description,
	})
	if err != nil {
		return err
	}

	return nil
}
