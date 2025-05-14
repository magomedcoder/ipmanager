package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/domain/entity"
	postgresModel "github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	postgresRepo "github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/repository"
	"github.com/magomedcoder/ipmanager/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
)

type ISubnetUseCase interface {
	Create(ctx context.Context, opt *SubnetOpt) (*entity.Subnet, error)

	GetSubnets(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Subnet, error)

	GetById(ctx context.Context, id int64) (*entity.Subnet, error)

	EditVlanById(ctx context.Context, id int64, vlanId int64) error

	EditDescriptionById(ctx context.Context, id int64, description string) error
}

var _ ISubnetUseCase = (*SubnetUseCase)(nil)

type SubnetUseCase struct {
	Conf       *config.Config
	SubnetRepo postgresRepo.ISubnetRepository
	IpRepo     postgresRepo.IIpRepository
}

type SubnetOpt struct {
	Id          int64
	Ip          string
	VlanId      *int64
	Description string
}

func (s *SubnetUseCase) Create(ctx context.Context, opt *SubnetOpt) (*entity.Subnet, error) {
	subnet, err := s.SubnetRepo.GetBySubnet(opt.Ip)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return nil, err
		}
	}

	if subnet != nil && subnet.ID != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Подсеть %s уже существует", subnet.Ip))
	}

	createdSubnet, err := s.SubnetRepo.Create(ctx, &postgresModel.Subnet{
		Ip:          opt.Ip,
		VlanID:      opt.VlanId,
		Description: opt.Description,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать подсеть")
	}

	ipPool, err := pkg.GenerateIPPool(opt.Ip)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать подсеть")
	}

	var ips []postgresModel.Ip
	for _, ip := range ipPool {
		ips = append(ips, postgresModel.Ip{
			SubnetID: createdSubnet.ID,
			Ip:       ip,
		})
	}

	if err := s.IpRepo.CreateBatches(ctx, ips); err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать подсеть")
	}

	return &entity.Subnet{
		Id: int64(createdSubnet.ID),
		Ip: createdSubnet.Ip,
	}, nil
}

func (s *SubnetUseCase) GetSubnets(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.Subnet, error) {
	subnets, err := s.SubnetRepo.GetSubnets(ctx, func(db *gorm.DB) {
		for _, fn := range arg {
			fn(db)
		}
		db.Preload("Vlan")
	})
	if err != nil {
		return nil, err
	}

	items := make([]*entity.Subnet, 0)
	for _, item := range subnets {
		res := &entity.Subnet{
			Id: int64(item.ID),
			Ip: item.Ip,
		}

		if item.Vlan.ID != 0 {
			res.VlanId = int64(item.Vlan.ID)
			res.VlanName = item.Vlan.Name
		}

		items = append(items, res)
	}

	return items, nil
}

func (s *SubnetUseCase) GetById(ctx context.Context, id int64) (*entity.Subnet, error) {
	subnet, err := s.SubnetRepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Не удалось получить Subnet: %v", id))
	}
	res := &entity.Subnet{
		Id:          int64(subnet.ID),
		Ip:          subnet.Ip,
		Description: subnet.Description,
	}

	if subnet.Vlan != nil {
		res.VlanId = int64(subnet.Vlan.ID)
		res.VlanName = subnet.Vlan.Name
	}

	return res, nil
}

func (s *SubnetUseCase) EditVlanById(ctx context.Context, id int64, vlanId int64) error {
	if err := s.SubnetRepo.EditVlanById(ctx, id, vlanId); err != nil {
		fmt.Println(err)
		return errors.New(fmt.Sprintf("Не удалось получить ip: %v", id))
	}

	return nil
}

func (s *SubnetUseCase) EditDescriptionById(ctx context.Context, id int64, description string) error {
	if err := s.SubnetRepo.EditDescriptionById(ctx, id, description); err != nil {
		return errors.New(fmt.Sprintf("Не удалось получить ip: %v", id))
	}

	return nil
}
