package handler

import (
	"context"
	"github.com/magomedcoder/ipmanager/api/pb"
	"github.com/magomedcoder/ipmanager/internal/domain/entity"
	"github.com/magomedcoder/ipmanager/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
