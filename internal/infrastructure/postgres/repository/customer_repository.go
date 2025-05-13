package repository

import (
	"context"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"gorm.io/gorm"
	"log"
)

type ICustomerRepository interface {
	Create(ctx context.Context, customer *model.Customer) (*model.Customer, error)

	GetCustomers(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Customer, error)

	GetById(ctx context.Context, id int64) (*model.Customer, error)
}

var _ ICustomerRepository = (*CustomerRepository)(nil)

type CustomerRepository struct {
	gormutil.Repo[model.Customer]
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{Repo: gormutil.NewRepo[model.Customer](db)}
}

func (c *CustomerRepository) Create(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	if err := c.Repo.Create(ctx, customer); err != nil {
		log.Printf("Не удалось создать пользователя: %s", err)
		return nil, err
	}

	return customer, nil
}

func (c *CustomerRepository) GetCustomers(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.Customer, error) {
	customers, err := c.FindAll(ctx, arg...)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *CustomerRepository) GetById(ctx context.Context, id int64) (*model.Customer, error) {
	customer, err := c.Repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
