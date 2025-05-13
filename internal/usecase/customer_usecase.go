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

type ICustomerUseCase interface {
	Create(ctx context.Context, opt *CustomerOpt) (*entity.Customer, error)

	GetCustomers(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Customer, error)

	GetById(ctx context.Context, id int64) (*entity.Customer, error)
}

var _ ICustomerUseCase = (*CustomerUseCase)(nil)

type CustomerUseCase struct {
	Conf         *config.Config
	CustomerRepo postgresRepo.ICustomerRepository
}

type CustomerOpt struct {
	Id      int64
	Name    string
	Surname string
}

func (c *CustomerUseCase) Create(ctx context.Context, opt *CustomerOpt) (*entity.Customer, error) {
	customer, err := c.CustomerRepo.Create(ctx, &postgresModel.Customer{
		Name:    opt.Name,
		Surname: opt.Surname,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать customer")
	}

	return &entity.Customer{
		Id:      int64(customer.ID),
		Name:    customer.Name,
		Surname: customer.Name,
	}, nil
}

func (c *CustomerUseCase) GetCustomers(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Customer, error) {
	customers, err := c.CustomerRepo.GetCustomers(ctx, arg...)
	if err != nil {
		return nil, err
	}

	items := make([]*entity.Customer, 0)
	for _, item := range customers {
		items = append(items, &entity.Customer{
			Id:      int64(item.ID),
			Name:    item.Name,
			Surname: item.Name,
		})
	}

	return items, nil
}

func (c *CustomerUseCase) GetById(ctx context.Context, id int64) (*entity.Customer, error) {
	item, err := c.CustomerRepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Не удалось получить customer: %v", id))
	}

	return &entity.Customer{
		Id:      int64(item.ID),
		Name:    item.Name,
		Surname: item.Name,
	}, nil
}
