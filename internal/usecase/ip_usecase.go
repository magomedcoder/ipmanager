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
	"log"
)

type IIpUseCase interface {
	Create(ctx context.Context, opt *IpOpt) (*entity.Ip, error)

	GetIps(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Ip, error)

	GetById(ctx context.Context, id int64) (*entity.Ip, error)
}

var _ IIpUseCase = (*IpUseCase)(nil)

type IpUseCase struct {
	Conf   *config.Config
	IpRepo postgresRepo.IIpRepository
}

type IpOpt struct {
	Id int64
	Ip string
}

func (i *IpUseCase) Create(ctx context.Context, opt *IpOpt) (*entity.Ip, error) {
	ip, err := i.IpRepo.GetByIp(opt.Ip)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return nil, err
		}
	}

	if ip != nil && ip.ID != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("IP %s уже существует", ip.Ip))
	}

	createdIp, err := i.IpRepo.Create(ctx, &postgresModel.Ip{
		Ip: opt.Ip,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать IP")
	}

	return &entity.Ip{
		Id: int64(createdIp.ID),
		Ip: createdIp.Ip,
	}, nil
}

func (i *IpUseCase) GetIps(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Ip, error) {
	ips, err := i.IpRepo.GetIps(ctx, func(db *gorm.DB) {
		for _, fn := range arg {
			fn(db)
		}
		db.Preload("Vlan").Preload("Customer")
	})
	if err != nil {
		return nil, err
	}

	items := make([]*entity.Ip, 0)
	for _, item := range ips {
		fmt.Println(item.Vlan)
		res := &entity.Ip{
			Id: int64(item.ID),
			Ip: item.Ip,
		}

		if item.Vlan.ID != 0 {
			res.VlanId = int64(item.Vlan.ID)
			res.VlanName = item.Vlan.Name
		}

		if item.Customer.ID != 0 {
			res.CustomerId = int64(item.Customer.ID)
			res.CustomerName = item.Customer.Name
		}
		items = append(items, res)
	}

	return items, nil
}

func (i *IpUseCase) GetById(ctx context.Context, id int64) (*entity.Ip, error) {
	ip, err := i.IpRepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Не удалось получить ip: %v", id))
	}

	return &entity.Ip{
		Id: int64(ip.ID),
		Ip: ip.Ip,
	}, nil
}
