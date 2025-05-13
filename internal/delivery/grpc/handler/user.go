package handler

import (
	"context"
	"github.com/magomedcoder/ipmanager/api/pb"
	"github.com/magomedcoder/ipmanager/internal/domain/entity"
	"github.com/magomedcoder/ipmanager/internal/usecase"
	"github.com/magomedcoder/ipmanager/pkg/gormutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	UserUseCase usecase.IUserUseCase
}

func (u *UserHandler) Login(ctx context.Context, in *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	token, err := u.UserUseCase.Login(ctx, in.Username, in.Password)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Ошибка создания токена")
	}

	return &pb.UserLoginResponse{
		AccessToken: token.AccessToken,
	}, nil
}

func (u *UserHandler) Logout(ctx context.Context, in *pb.UserLogoutRequest) (*pb.UserLogoutResponse, error) {
	claims, ok := entity.GetUserClaimsFromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "не удалось получить информацию")
	}

	if err := u.UserUseCase.Logout(ctx, claims.Token); err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Ошибка создания токена")
	}

	return &pb.UserLogoutResponse{
		Success: true,
	}, nil
}

func (u *UserHandler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userOpt := &usecase.UserOpt{
		Username: in.Username,
		Password: in.Password,
		Name:     in.Name,
		Surname:  in.Surname,
	}

	user, err := u.UserUseCase.Create(ctx, userOpt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Ошибка")
	}

	return &pb.CreateUserResponse{
		Id: user.Id,
	}, nil
}

func (u *UserHandler) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	pagination := &gormutil.Pagination{
		Page:     1,
		PageSize: 15,
	}
	pagination.SetPage(int(in.GetPage()))
	pagination.SetPageSize(int(in.GetPageSize()))

	var total int64
	users, err := u.UserUseCase.GetUsers(ctx, func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(pagination)).
			Count(&total)
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "Пользователи не найдены")
	}

	items := make([]*pb.UserItem, 0)
	for _, item := range users {
		items = append(items, &pb.UserItem{
			Id:       item.Id,
			Username: item.Username,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}

	return &pb.GetUsersResponse{
		Total: total,
		Items: items,
	}, nil
}

func (u *UserHandler) GetUserById(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, _ := u.UserUseCase.GetUserById(ctx, in.Id)
	if user.Id == 0 {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	return &pb.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
	}, nil
}
