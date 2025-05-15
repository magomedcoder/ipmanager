package repository

import (
	"context"
	"errors"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"gorm.io/gorm"
	"log"
)

type IIpRepository interface {
	Create(ctx context.Context, ip *model.Ip) (*model.Ip, error)

	CreateBatches(ctx context.Context, ip []model.Ip) error

	GetIps(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Ip, error)

	GetById(ctx context.Context, id int64) (*model.Ip, error)

	GetByIp(ip string) (*model.Ip, error)

	EditCustomerById(ctx context.Context, id int64, customerId int64) error

	EditServiceById(ctx context.Context, id int64, serviceId int64) error

	EditDescriptionById(ctx context.Context, id int64, description string) error

	CountByCustomerBinding(ctx context.Context, subnetID int64) (int64, int64, error)
}

var _ IIpRepository = (*IpRepository)(nil)

type IpRepository struct {
	gormutil.Repo[model.Ip]
}

func NewIPRepository(db *gorm.DB) *IpRepository {
	return &IpRepository{Repo: gormutil.NewRepo[model.Ip](db)}
}

func (i *IpRepository) Create(ctx context.Context, ip *model.Ip) (*model.Ip, error) {
	if err := i.Repo.Create(ctx, ip); err != nil {
		log.Printf("Не удалось создать ip: %s", err)
		return nil, err
	}

	return ip, nil
}

func (i *IpRepository) CreateBatches(ctx context.Context, ip []model.Ip) error {
	if err := i.Repo.CreateInBatches(ctx, ip, 225); err != nil {
		log.Printf("Не удалось создать ip: %s", err)
		return err
	}

	return nil
}

func (i *IpRepository) GetIps(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Ip, error) {
	ips, err := i.FindAll(ctx, arg...)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func (i *IpRepository) GetById(ctx context.Context, id int64) (*model.Ip, error) {
	ip, err := i.Repo.FindByWhereWithQuery(ctx, "id = ?", []any{id}, func(db *gorm.DB) {
		db.Preload("Customer").Preload("Service")
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Не удалось получить: %s", err)
		}
		return nil, err
	}

	return ip, nil
}

func (i *IpRepository) GetByIp(ip string) (*model.Ip, error) {
	ipData, err := i.Repo.FindByWhere(context.TODO(), "ip = ?", ip)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Не удалось получить пользователя: %s", err)
		}
		return nil, err
	}

	return ipData, nil
}

func (i *IpRepository) EditCustomerById(ctx context.Context, id int64, customerId int64) error {
	data := map[string]any{
		"customer_id": nil,
	}

	if customerId != 0 {
		data["customer_id"] = customerId
	}

	_, err := i.Repo.UpdateById(ctx, id, data)
	if err != nil {
		return err
	}

	return nil
}

func (i *IpRepository) EditServiceById(ctx context.Context, id int64, serviceId int64) error {
	data := map[string]any{
		"service_id": nil,
	}

	if serviceId != 0 {
		data["service_id"] = serviceId
	}

	_, err := i.Repo.UpdateById(ctx, id, data)
	if err != nil {
		return err
	}

	return nil
}

func (i *IpRepository) EditDescriptionById(ctx context.Context, id int64, description string) error {
	_, err := i.Repo.UpdateById(ctx, id, map[string]any{
		"description": description,
	})
	if err != nil {
		return err
	}

	return nil
}

func (i *IpRepository) CountByCustomerBinding(ctx context.Context, subnetID int64) (int64, int64, error) {
	var withoutCustomer int64
	var withCustomer int64

	if err := i.Repo.Model(ctx).
		Where("subnet_id = ? AND customer_id IS NULL", subnetID).
		Count(&withoutCustomer).Error; err != nil {
		return 0, 0, err
	}

	if err := i.Repo.Model(ctx).
		Where("subnet_id = ? AND customer_id IS NOT NULL", subnetID).
		Count(&withCustomer).Error; err != nil {
		return 0, 0, err
	}

	return withoutCustomer, withCustomer, nil
}
