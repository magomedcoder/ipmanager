package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/domain/entity"
	postgresModel "github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/model"
	postgresRepo "github.com/magomedcoder/ipmanager/internal/infrastructure/postgres/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"time"
)

type IUserUseCase interface {
	Login(ctx context.Context, username string, password string) (*entity.UserLogin, error)

	Logout(ctx context.Context, accessToken string) error

	Create(ctx context.Context, userModel *postgresModel.User) (*postgresModel.User, error)

	ValidateToken(tokenString string) (*UserClaims, error)

	IsTokenRevoked(ctx context.Context, token string) (bool, error)
}

var _ IUserUseCase = (*UserUseCase)(nil)

type UserUseCase struct {
	Conf            *config.Config
	UserRepo        postgresRepo.IUserRepository
	UserSessionRepo postgresRepo.IUserSessionRepository
}

func (u *UserUseCase) Login(ctx context.Context, username string, password string) (*entity.UserLogin, error) {
	user, err := u.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	if _, err := u.checkPasswordHash(password, user.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный пароль")
	}

	token, err := u.generateToken(user)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "ошибка создания токена")
	}

	if _, err := u.UserSessionRepo.Create(ctx, &postgresModel.UserSession{
		AccessToken: token.AccessToken,
		ExpiresAt:   token.ExpiresAt,
		UserId:      user.ID,
	}); err != nil {
		return nil, status.Error(codes.FailedPrecondition, "ошибка создания токена")
	}

	return token, nil
}

func (u *UserUseCase) Logout(ctx context.Context, accessToken string) error {
	return u.UserSessionRepo.DeleteByToken(ctx, accessToken)
}

func (u *UserUseCase) Create(ctx context.Context, userModel *postgresModel.User) (*postgresModel.User, error) {
	user, err := u.UserRepo.GetByUsername(userModel.Username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return nil, err
		}
	}

	if user != nil && user.ID != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Пользователь %s уже существует", user.Username))
	}

	passwordHash, err := u.hashPassword(userModel.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось хешировать пароль")
	}

	userModel.Password = passwordHash

	createdUser, err := u.UserRepo.Create(ctx, userModel)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	return createdUser, nil
}

func (u *UserUseCase) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *UserUseCase) checkPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

type UserClaims struct {
	*jwt.StandardClaims
	entity.UserClaims
}

func (u *UserUseCase) generateToken(user *postgresModel.User) (*entity.UserLogin, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(u.Conf.Jwt.ExpiresTime)).Unix()
	claims := &UserClaims{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		entity.UserClaims{
			UId:      user.ID,
			Username: user.Username,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(u.Conf.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	return &entity.UserLogin{
		AccessToken: token,
		ExpiresAt:   expiresAt,
	}, nil
}

func (u *UserUseCase) ValidateToken(tokenString string) (*UserClaims, error) {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.Conf.Jwt.Secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("недействительный токен")
	}

	expirationTime := time.Unix(claims.ExpiresAt, 0)
	if expirationTime.Before(time.Now()) {
		return nil, errors.New("токен аннулирован")
	}

	return claims, nil
}

func (u *UserUseCase) IsTokenRevoked(ctx context.Context, accessToken string) (bool, error) {
	isValid, err := u.UserSessionRepo.CheckExpiredToken(ctx, accessToken)
	if err != nil {
		return true, err
	}

	return !isValid, nil
}
