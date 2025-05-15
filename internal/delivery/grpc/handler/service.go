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

type ServiceHandler struct {
	pb.UnimplementedServiceServiceServer
	ServiceUseCase usecase.IServiceUseCase
}

func (v *ServiceHandler) CreateService(ctx context.Context, in *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	serviceOpt := &usecase.ServiceOpt{
		Name: in.Name,
	}

	service, err := v.ServiceUseCase.Create(ctx, serviceOpt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Ошибка")
	}

	return &pb.CreateServiceResponse{
		Id: service.Id,
	}, nil
}

func (v *ServiceHandler) GetServices(ctx context.Context, in *pb.GetServicesRequest) (*pb.GetServicesResponse, error) {
	pagination := &gormutil.Pagination{
		Page:     1,
		PageSize: 15,
	}
	pagination.SetPage(int(in.GetPage()))
	pagination.SetPageSize(int(in.GetPageSize()))

	var total int64
	services, err := v.ServiceUseCase.GetServices(ctx, func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(pagination)).
			Count(&total)
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "Service не найдены")
	}

	items := make([]*pb.ServiceItem, 0)
	for _, item := range services {
		var vlans []*pb.ServiceVlan
		for _, _vlan := range item.Vlans {
			vlans = append(vlans, &pb.ServiceVlan{
				Id:   _vlan.Id,
				Name: _vlan.Name,
			})
		}

		var ips []*pb.ServiceIp
		for _, _ip := range item.Ips {
			ips = append(ips, &pb.ServiceIp{
				Id: _ip.Id,
				Ip: _ip.Ip,
			})
		}

		items = append(items, &pb.ServiceItem{
			Id:    item.Id,
			Name:  item.Name,
			Vlans: vlans,
			Ips:   ips,
		})
	}

	return &pb.GetServicesResponse{
		Total: total,
		Items: items,
	}, nil
}

func (v *ServiceHandler) GetServiceById(ctx context.Context, in *pb.GetServiceRequest) (*pb.GetServiceResponse, error) {
	service, _ := v.ServiceUseCase.GetById(ctx, in.Id)
	if service.Id == 0 {
		return nil, status.Error(codes.NotFound, "Service не найден")
	}

	var ips []*pb.ServiceIp
	for _, _ip := range service.Ips {
		ips = append(ips, &pb.ServiceIp{
			Id: _ip.Id,
			Ip: _ip.Ip,
		})
	}

	return &pb.GetServiceResponse{
		Id:   service.Id,
		Name: service.Name,
		Ips:  ips,
	}, nil
}
