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

	EditCustomerById(ctx context.Context, id int64, customerId int64) error

	EditServiceById(ctx context.Context, id int64, serviceId int64) error

	EditDescriptionById(ctx context.Context, id int64, description string) error

	CountByCustomerBinding(ctx context.Context, id int64) (int64, int64, error)
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
		db.Preload("Customer")
	})
	if err != nil {
		return nil, err
	}

	items := make([]*entity.Ip, 0)
	for _, item := range ips {
		res := &entity.Ip{
			Id:          int64(item.ID),
			Ip:          item.Ip,
			SubnetId:    int64(item.SubnetID),
			SubnetName:  item.Ip,
			Description: item.Description,
		}

		if item.Customer != nil {
			res.Busy = true
		}

		if item.Customer != nil {
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

	res := &entity.Ip{
		Id:          int64(ip.ID),
		Ip:          ip.Ip,
		SubnetId:    int64(ip.SubnetID),
		SubnetName:  ip.Ip,
		Description: ip.Description,
	}

	if ip.Customer != nil {
		res.Busy = true
	}

	if ip.Customer != nil {
		res.CustomerId = int64(ip.Customer.ID)
		res.CustomerName = ip.Customer.Name
	}

	return res, nil
}

func (i *IpUseCase) EditCustomerById(ctx context.Context, id int64, customerId int64) error {
	if err := i.IpRepo.EditCustomerById(ctx, id, customerId); err != nil {
		return errors.New(fmt.Sprintf("Не удалось получить ip: %v", id))
	}

	return nil
}

func (i *IpUseCase) EditServiceById(ctx context.Context, id int64, serviceId int64) error {
	//if err := i.IpRepo.EditServiceById(ctx, id, serviceId); err != nil {
	//	return errors.New(fmt.Sprintf("Не удалось получить ip: %v", id))
	//}

	return nil
}

func (i *IpUseCase) EditDescriptionById(ctx context.Context, id int64, description string) error {
	if err := i.IpRepo.EditDescriptionById(ctx, id, description); err != nil {
		return errors.New(fmt.Sprintf("Не удалось получить ip: %v", id))
	}

	return nil
}

func (i *IpUseCase) CountByCustomerBinding(ctx context.Context, id int64) (int64, int64, error) {
	withoutCustomer, withCustomer, err := i.IpRepo.CountByCustomerBinding(ctx, id)
	if err != nil {
		return 0, 0, status.Error(codes.Internal, "Ошибка при подсчете IP-адресов")
	}

	return withoutCustomer, withCustomer, nil
}
