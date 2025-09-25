package di

import (
	"github.com/magomedcoder/ipmanager/api/pb"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/handler"
	"github.com/magomedcoder/ipmanager/internal/delivery/grpc/middleware"
	"github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/repository"
	"github.com/magomedcoder/ipmanager/internal/provider"
	"github.com/magomedcoder/ipmanager/internal/usecase"
)

type GrpcProvider struct {
	Middleware      middleware.AuthMiddleware
	UserHandler     pb.UserServiceServer
	SubnetHandler   pb.SubnetServiceServer
	IpHandler       pb.IpServiceServer
	CustomerHandler pb.CustomerServiceServer
	VlanHandler     pb.VlanServiceServer
	ServiceHandler  pb.ServiceServiceServer
}

func NewGrpcInjector(conf *config.Config) *GrpcProvider {
	db := provider.NewPostgresClient(conf)
	userRepository := repository.NewUserRepository(db)
	userSessionRepository := repository.NewUserSessionRepository(db)
	userUseCase := &usecase.UserUseCase{
		Conf:            conf,
		UserRepo:        userRepository,
		UserSessionRepo: userSessionRepository,
	}
	authMiddleware := middleware.AuthMiddleware{
		UserUseCase: userUseCase,
	}
	unimplementedUserServiceServer := pb.UnimplementedUserServiceServer{}
	userHandler := &handler.UserHandler{
		UnimplementedUserServiceServer: unimplementedUserServiceServer,
		UserUseCase:                    userUseCase,
	}
	unimplementedSubnetServiceServer := pb.UnimplementedSubnetServiceServer{}
	subnetRepository := repository.NewSubnetRepository(db)
	ipRepository := repository.NewIPRepository(db)
	subnetUseCase := &usecase.SubnetUseCase{
		Conf:       conf,
		SubnetRepo: subnetRepository,
		IpRepo:     ipRepository,
	}
	ipUseCase := &usecase.IpUseCase{
		Conf:   conf,
		IpRepo: ipRepository,
	}
	subnetHandler := &handler.SubnetHandler{
		UnimplementedSubnetServiceServer: unimplementedSubnetServiceServer,
		SubnetUseCase:                    subnetUseCase,
		IpUseCase:                        ipUseCase,
	}
	unimplementedIpServiceServer := pb.UnimplementedIpServiceServer{}
	ipHandler := &handler.IpHandler{
		UnimplementedIpServiceServer: unimplementedIpServiceServer,
		IpUseCase:                    ipUseCase,
		SubnetUseCase:                subnetUseCase,
	}
	unimplementedCustomerServiceServer := pb.UnimplementedCustomerServiceServer{}
	customerRepository := repository.NewCustomerRepository(db)
	customerUseCase := &usecase.CustomerUseCase{
		Conf:         conf,
		CustomerRepo: customerRepository,
	}
	customerHandler := &handler.CustomerHandler{
		UnimplementedCustomerServiceServer: unimplementedCustomerServiceServer,
		CustomerUseCase:                    customerUseCase,
	}
	unimplementedVlanServiceServer := pb.UnimplementedVlanServiceServer{}
	vlanRepository := repository.NewVlanRepository(db)
	vlanUseCase := &usecase.VlanUseCase{
		Conf:     conf,
		VlanRepo: vlanRepository,
	}
	vlanHandler := &handler.VlanHandler{
		UnimplementedVlanServiceServer: unimplementedVlanServiceServer,
		VlanUseCase:                    vlanUseCase,
	}
	unimplementedServiceServiceServer := pb.UnimplementedServiceServiceServer{}
	serviceRepository := repository.NewServiceRepository(db)
	serviceUseCase := &usecase.ServiceUseCase{
		Conf:        conf,
		ServiceRepo: serviceRepository,
	}
	serviceHandler := &handler.ServiceHandler{
		UnimplementedServiceServiceServer: unimplementedServiceServiceServer,
		ServiceUseCase:                    serviceUseCase,
	}
	appProvider := &GrpcProvider{
		Middleware:      authMiddleware,
		UserHandler:     userHandler,
		SubnetHandler:   subnetHandler,
		IpHandler:       ipHandler,
		CustomerHandler: customerHandler,
		VlanHandler:     vlanHandler,
		ServiceHandler:  serviceHandler,
	}

	return appProvider
}
