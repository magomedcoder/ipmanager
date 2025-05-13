package middleware

import (
	"context"
	"github.com/magomedcoder/ipmanager/internal/domain/entity"
	"github.com/magomedcoder/ipmanager/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "отсутствуют метаданные")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "недействительный токен")
	errRevokedToken    = status.Errorf(codes.Unauthenticated, "токен аннулирован")
)

type AuthMiddleware struct {
	UserUseCase usecase.IUserUseCase
}

func (a *AuthMiddleware) validateToken(ctx context.Context) (*entity.UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	tokenStr := ""
	if authHeader, exists := md["authorization"]; exists && len(authHeader) > 0 {
		tokenStr = authHeader[0]
	} else {
		return nil, errInvalidToken
	}

	token := strings.TrimPrefix(tokenStr, "Bearer ")

	claims, err := a.UserUseCase.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	if revoked, _ := a.UserUseCase.IsTokenRevoked(ctx, token); revoked {
		return nil, errRevokedToken
	}

	return &entity.UserClaims{
		UId:      claims.UId,
		Username: claims.Username,
		Token:    token,
	}, nil
}

func (a *AuthMiddleware) UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if info.FullMethod == "/user.UserService/Login" {
		return handler(ctx, req)
	}

	claims, err := a.validateToken(ctx)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, entity.UserClaimsContextKey, claims)

	return handler(ctx, req)
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}

func (a *AuthMiddleware) StreamInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	claims, err := a.validateToken(ss.Context())
	if err != nil {
		return err
	}

	ctx := context.WithValue(ss.Context(), entity.UserClaimsContextKey, claims)

	wrapped := &wrappedStream{ServerStream: ss, ctx: ctx}
	return handler(srv, wrapped)
}
