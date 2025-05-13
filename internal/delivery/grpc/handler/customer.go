package handler

import (
	"context"
	"github.com/magomedcoder/ipmanager/api/pb"
	"github.com/magomedcoder/ipmanager/internal/usecase"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	CustomerUseCase usecase.ICustomerUseCase
}

func (c *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	customerOpt := &usecase.CustomerOpt{
		Name:    in.Name,
		Surname: in.Surname,
	}

	customer, err := c.CustomerUseCase.Create(ctx, customerOpt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Ошибка")
	}

	return &pb.CreateCustomerResponse{
		Id: customer.Id,
	}, nil
}

func (c *CustomerHandler) GetCustomers(ctx context.Context, in *pb.GetCustomersRequest) (*pb.GetCustomersResponse, error) {
	pagination := &gormutil.Pagination{
		Page:     1,
		PageSize: 15,
	}
	pagination.SetPage(int(in.GetPage()))
	pagination.SetPageSize(int(in.GetPageSize()))

	var total int64
	customers, err := c.CustomerUseCase.GetCustomers(ctx, func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(pagination)).
			Count(&total)
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "Customer не найдены")
	}

	items := make([]*pb.CustomerItem, 0)
	for _, item := range customers {
		items = append(items, &pb.CustomerItem{
			Id:      item.Id,
			Name:    item.Name,
			Surname: item.Surname,
		})
	}

	return &pb.GetCustomersResponse{
		Total: total,
		Items: items,
	}, nil
}

func (c *CustomerHandler) GetCustomerById(ctx context.Context, in *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	customer, _ := c.CustomerUseCase.GetById(ctx, in.Id)
	if customer.Id == 0 {
		return nil, status.Error(codes.NotFound, "Customer не найден")
	}

	return &pb.GetCustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		Surname: customer.Surname,
	}, nil
}
