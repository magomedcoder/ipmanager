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

	ChangePasswordById(ctx context.Context, id int64, password string, newPassword string) error

	ValidateToken(tokenString string) (*UserClaims, error)

	IsTokenRevoked(ctx context.Context, token string) (bool, error)

	Create(ctx context.Context, opt *UserOpt) (*entity.User, error)

	GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.User, error)

	GetUserById(ctx context.Context, id int64) (*entity.User, error)
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
			UId:      int64(user.ID),
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

func (u *UserUseCase) ChangePasswordById(ctx context.Context, id int64, password string, newPassword string) error {
	user, err := u.UserRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if _, err := u.checkPasswordHash(password, user.Password); err != nil {
		return errors.New("Неверно введен пароль")
	}

	passwordHash, err := u.hashPassword(newPassword)
	if err != nil {
		return errors.New("Не удалось хешировать пароль")
	}

	_, err = u.UserRepo.UpdatePasswordById(ctx, id, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

type UserOpt struct {
	Id       int64
	Username string
	Password string
	Name     string
	Surname  string
}

func (u *UserUseCase) Create(ctx context.Context, opt *UserOpt) (*entity.User, error) {
	user, err := u.UserRepo.GetByUsername(opt.Username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
			return nil, err
		}
	}

	if user != nil && user.ID != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Пользователь %s уже существует", user.Username))
	}

	passwordHash, err := u.hashPassword(opt.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось хешировать пароль")
	}

	createdUser, err := u.UserRepo.Create(ctx, &postgresModel.User{
		Username: opt.Username,
		Password: passwordHash,
		Name:     opt.Name,
		Surname:  opt.Surname,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	return &entity.User{
		Id:       int64(createdUser.ID),
		Username: createdUser.Username,
		Name:     createdUser.Name,
		Surname:  createdUser.Surname,
	}, nil
}

func (u *UserUseCase) GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.User, error) {
	users, err := u.UserRepo.GetUsers(ctx, arg...)
	if err != nil {
		return nil, err
	}

	items := make([]*entity.User, 0)
	for _, item := range users {
		items = append(items, &entity.User{
			Id:       int64(item.ID),
			Username: item.Username,
			Name:     item.Name,
			Surname:  item.Surname,
		})
	}

	return items, nil
}

func (u *UserUseCase) GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	user, err := u.UserRepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Не удалось получить пользователя: %v", id))
	}

	return &entity.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
	}, nil
}
