package handler

import (
	"context"
	"github.com/magomedcoder/ipmanager/api/pb"
	"github.com/magomedcoder/ipmanager/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type IpHandler struct {
	pb.UnimplementedIpServiceServer
	IpUseCase     usecase.IIpUseCase
	SubnetUseCase usecase.ISubnetUseCase
}

func (i *IpHandler) GetIps(ctx context.Context, in *pb.GetIpsRequest) (*pb.GetIpsResponse, error) {
	var total int64
	ips, err := i.IpUseCase.GetIps(ctx, func(db *gorm.DB) {
		db.Where("subnet_id = ?", in.SubnetId).
			Order("id ASC").
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
			Description:  item.Description,
		})
	}

	return &pb.GetIpsResponse{
		Total: total,
		Items: items,
	}, nil
}

func (i *IpHandler) GetIpById(ctx context.Context, in *pb.GetIpRequest) (*pb.GetIpResponse, error) {
	ip, _ := i.IpUseCase.GetById(ctx, in.Id)
	if ip == nil {
		return nil, status.Error(codes.NotFound, "Ip не найден")
	}

	return &pb.GetIpResponse{
		Id:           ip.Id,
		Ip:           ip.Ip,
		CustomerId:   ip.CustomerId,
		CustomerName: ip.CustomerName,
		Description:  ip.Description,
	}, nil
}

func (i *IpHandler) EditIpCustomer(ctx context.Context, in *pb.EditIpCustomerRequest) (*pb.EditIpCustomerResponse, error) {
	if err := i.IpUseCase.EditCustomerById(ctx, in.Id, in.CustomerId); err != nil {
		return nil, status.Error(codes.NotFound, "Ip не найден")
	}
	return &pb.EditIpCustomerResponse{
		Success: true,
	}, nil
}

func (i *IpHandler) EditIpService(ctx context.Context, in *pb.EditIpServiceRequest) (*pb.EditIpServiceResponse, error) {
	if err := i.IpUseCase.EditServiceById(ctx, in.Id, in.ServiceId); err != nil {
		return nil, status.Error(codes.NotFound, "Ip не найден")
	}
	return &pb.EditIpServiceResponse{
		Success: true,
	}, nil
}

func (i *IpHandler) EditIpDescription(ctx context.Context, in *pb.EditIpDescriptionRequest) (*pb.EditIpDescriptionResponse, error) {
	if err := i.IpUseCase.EditDescriptionById(ctx, in.Id, in.Description); err != nil {
		return nil, status.Error(codes.NotFound, "Ip не найден")
	}

	return &pb.EditIpDescriptionResponse{
		Success: true,
	}, nil
}
