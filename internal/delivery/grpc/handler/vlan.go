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

type VlanHandler struct {
	pb.UnimplementedVlanServiceServer
	VlanUseCase usecase.IVlanUseCase
}

func (v *VlanHandler) CreateVlan(ctx context.Context, in *pb.CreateVlanRequest) (*pb.CreateVlanResponse, error) {
	vlanOpt := &usecase.VlanOpt{
		Name: in.Name,
	}

	vlan, err := v.VlanUseCase.Create(ctx, vlanOpt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Ошибка")
	}

	return &pb.CreateVlanResponse{
		Id: vlan.Id,
	}, nil
}

func (v *VlanHandler) GetVlans(ctx context.Context, in *pb.GetVlansRequest) (*pb.GetVlansResponse, error) {
	pagination := &gormutil.Pagination{
		Page:     1,
		PageSize: 15,
	}
	pagination.SetPage(int(in.GetPage()))
	pagination.SetPageSize(int(in.GetPageSize()))

	var total int64
	vlans, err := v.VlanUseCase.GetVlans(ctx, func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(pagination)).
			Count(&total)
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "Vlan не найдены")
	}

	items := make([]*pb.VlanItem, 0)
	for _, item := range vlans {
		items = append(items, &pb.VlanItem{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return &pb.GetVlansResponse{
		Total: total,
		Items: items,
	}, nil
}

func (v *VlanHandler) GetVlanById(ctx context.Context, in *pb.GetVlanRequest) (*pb.GetVlanResponse, error) {
	vlan, _ := v.VlanUseCase.GetById(ctx, in.Id)
	if vlan.Id == 0 {
		return nil, status.Error(codes.NotFound, "Vlan не найден")
	}

	return &pb.GetVlanResponse{
		Id:   vlan.Id,
		Name: vlan.Name,
	}, nil
}
