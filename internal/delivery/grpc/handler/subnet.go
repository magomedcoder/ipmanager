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

type SubnetHandler struct {
	pb.UnimplementedSubnetServiceServer
	SubnetUseCase usecase.ISubnetUseCase
}

func (s *SubnetHandler) CreateSubnet(ctx context.Context, in *pb.CreateSubnetRequest) (*pb.CreateSubnetResponse, error) {
	subnetOpt := &usecase.SubnetOpt{
		Ip:          in.Ip,
		Description: in.Description,
	}

	if in.VlanId != 0 {
		subnetOpt.VlanId = &in.VlanId
	}

	subnet, err := s.SubnetUseCase.Create(ctx, subnetOpt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Ошибка")
	}

	return &pb.CreateSubnetResponse{
		Id: subnet.Id,
	}, nil
}

func (s *SubnetHandler) GetSubnets(ctx context.Context, in *pb.GetSubnetsRequest) (*pb.GetSubnetsResponse, error) {
	pagination := &gormutil.Pagination{
		Page:     1,
		PageSize: 15,
	}
	pagination.SetPage(int(in.GetPage()))
	pagination.SetPageSize(int(in.GetPageSize()))

	var total int64
	subnets, err := s.SubnetUseCase.GetSubnets(ctx, func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(pagination)).
			Count(&total)
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "Subnet не найдены")
	}

	items := make([]*pb.SubnetItem, 0)
	for _, item := range subnets {
		items = append(items, &pb.SubnetItem{
			Id:           item.Id,
			Ip:           item.Ip,
			CustomerId:   item.CustomerId,
			CustomerName: item.CustomerName,
			VlanId:       item.VlanId,
			VlanName:     item.VlanName,
		})
	}

	return &pb.GetSubnetsResponse{
		Total: total,
		Items: items,
	}, nil
}

func (s *SubnetHandler) GetSubnetById(ctx context.Context, in *pb.GetSubnetRequest) (*pb.GetSubnetResponse, error) {
	subnet, _ := s.SubnetUseCase.GetById(ctx, in.Id)
	if subnet == nil {
		return nil, status.Error(codes.NotFound, "Подсеть не найден")
	}

	var charts []int64
	///
	charts = append(charts, 40)
	charts = append(charts, 60)
	//

	return &pb.GetSubnetResponse{
		Id:          subnet.Id,
		Ip:          subnet.Ip,
		VlanId:      subnet.VlanId,
		VlanName:    subnet.VlanName,
		Description: subnet.Description,
		Charts:      charts,
	}, nil
}

func (s *SubnetHandler) EditSubnetVlan(ctx context.Context, in *pb.EditSubnetVlanRequest) (*pb.EditSubnetVlanResponse, error) {
	if err := s.SubnetUseCase.EditVlanById(ctx, in.Id, in.VlanId); err != nil {
		return nil, status.Error(codes.NotFound, "Подсеть не найден")
	}

	return &pb.EditSubnetVlanResponse{
		Success: true,
	}, nil
}

func (s *SubnetHandler) EditSubnetDescription(ctx context.Context, in *pb.EditSubnetDescriptionRequest) (*pb.EditSubnetDescriptionResponse, error) {
	if err := s.SubnetUseCase.EditDescriptionById(ctx, in.Id, in.Description); err != nil {
		return nil, status.Error(codes.NotFound, "Подсеть не найден")
	}

	return &pb.EditSubnetDescriptionResponse{
		Success: true,
	}, nil
}
