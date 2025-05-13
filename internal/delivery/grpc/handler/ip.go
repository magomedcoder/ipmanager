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

type IpHandler struct {
	pb.UnimplementedIpServiceServer
	IpUseCase usecase.IIpUseCase
}

func (i *IpHandler) CreateIp(ctx context.Context, in *pb.CreateIpRequest) (*pb.CreateIpResponse, error) {
	ipOpt := &usecase.IpOpt{
		Ip: in.Ip,
	}

	ip, err := i.IpUseCase.Create(ctx, ipOpt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Ошибка")
	}

	return &pb.CreateIpResponse{
		Id: ip.Id,
	}, nil
}

func (i *IpHandler) GetIps(ctx context.Context, in *pb.GetIpsRequest) (*pb.GetIpsResponse, error) {
	pagination := &gormutil.Pagination{
		Page:     1,
		PageSize: 15,
	}
	pagination.SetPage(int(in.GetPage()))
	pagination.SetPageSize(int(in.GetPageSize()))

	var total int64
	ips, err := i.IpUseCase.GetIps(ctx, func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(pagination)).
			Count(&total)
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "IP не найдены")
	}

	items := make([]*pb.IpItem, 0)
	for _, item := range ips {
		items = append(items, &pb.IpItem{
			Id:           item.Id,
			Ip:           item.Ip,
			CustomerId:   item.CustomerId,
			CustomerName: item.CustomerName,
			VlanId:       item.VlanId,
			VlanName:     item.VlanName,
		})
	}

	return &pb.GetIpsResponse{
		Total: total,
		Items: items,
	}, nil
}

func (i *IpHandler) GetIpById(ctx context.Context, in *pb.GetIpRequest) (*pb.GetIpResponse, error) {
	ip, _ := i.IpUseCase.GetById(ctx, in.Id)
	if ip.Id == 0 {
		return nil, status.Error(codes.NotFound, "Ip не найден")
	}

	return &pb.GetIpResponse{
		Id: ip.Id,
		Ip: ip.Ip,
	}, nil
}
