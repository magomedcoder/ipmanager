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

	GetIps(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Ip, error)

	Get(ctx context.Context, id int64) (*model.Ip, error)

	GetById(ip string) (*model.Ip, error)
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

func (i *IpRepository) GetIps(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Ip, error) {
	ips, err := i.FindAll(ctx, arg...)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func (i *IpRepository) Get(ctx context.Context, id int64) (*model.Ip, error) {
	ip, err := i.Repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return ip, nil
}

func (i *IpRepository) GetById(ip string) (*model.Ip, error) {
	ipData, err := i.Repo.FindByWhere(context.TODO(), "ip = ?", ip)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Не удалось получить пользователя: %s", err)
		}
		return nil, err
	}

	return ipData, nil
}
